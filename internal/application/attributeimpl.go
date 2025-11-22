package application

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/constant"
	"backend/internal/domain"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type AttributeImpl struct {
	attributeRepo    domain.AttributeRepository
	attributeService domain.AttributeService
	redisClient      *redis.Client
}

func ProvideAttribute(attributeRepo domain.AttributeRepository, attributeService domain.AttributeService, redisClient *redis.Client) *AttributeImpl {
	return &AttributeImpl{
		attributeRepo:    attributeRepo,
		attributeService: attributeService,
		redisClient:      redisClient,
	}
}

var _ Attribute = &AttributeImpl{}

func (a *AttributeImpl) Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	attribute, err := a.attributeService.Create(param.Data.Code, param.Data.Name)
	if err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	// Invalidate list cache
	if a.redisClient != nil {
		iter := a.redisClient.Scan(ctx, 0, constant.AttributeListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
	}

	return attribute, nil
}

func (a *AttributeImpl) CreateValue(ctx context.Context, param CreateAttributeValueParam) (*domain.AttributeValue, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	attributeValue, err := a.attributeService.CreateValue(param.Data.Value)
	if err != nil {
		return nil, err
	}
	err = a.attributeService.AddValues(*attribute, *attributeValue)
	if err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	if a.redisClient != nil {
		// Delete attribute cache (since it contains values)
		a.redisClient.Del(ctx, constant.AttributeGetKey(param.AttributeID))
		// Delete attribute list cache
		iter := a.redisClient.Scan(ctx, 0, constant.AttributeListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
		// Delete attribute value list cache
		iter = a.redisClient.Scan(ctx, 0, constant.AttributeValueListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
	}

	return attributeValue, nil
}

func (a *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error) {
	// Build cache key
	var ids *[]uuid.UUID
	if param.AttributeIDs != nil {
		ids = param.AttributeIDs
	}
	searchStr := ""
	if param.Search != nil {
		searchStr = *param.Search
	}
	cacheKey := constant.AttributeListKey(ids, searchStr, string(param.Deleted), *param.Limit, *param.Page)

	// Try to get from cache
	if a.redisClient != nil {
		cachedData, err := a.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var pagination Pagination[domain.Attribute]
			if err := json.Unmarshal([]byte(cachedData), &pagination); err == nil {
				return &pagination, nil
			}
		}
	}

	attributes, err := a.attributeRepo.List(
		ctx,
		param.AttributeIDs,
		param.Search,
		param.Deleted,
		*param.Limit,
		*param.Page,
	)
	if err != nil {
		return nil, err
	}
	count, err := a.attributeRepo.Count(
		ctx,
		param.AttributeIDs,
		param.Deleted,
	)
	if err != nil {
		return nil, err
	}
	pagination := newPagination(*attributes, *count, *param.Page, *param.Limit)

	// Cache the result
	if a.redisClient != nil {
		if data, err := json.Marshal(pagination); err == nil {
			a.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLAttribute)*time.Second)
		}
	}

	return pagination, nil
}

func (a *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	// Build cache key
	cacheKey := constant.AttributeGetKey(param.AttributeID)

	// Try to get from cache
	if a.redisClient != nil {
		cachedData, err := a.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var attribute domain.Attribute
			if err := json.Unmarshal([]byte(cachedData), &attribute); err == nil {
				return &attribute, nil
			}
		}
	}

	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	if a.redisClient != nil {
		if data, err := json.Marshal(attribute); err == nil {
			a.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLAttribute)*time.Second)
		}
	}

	return attribute, nil
}

func (a *AttributeImpl) ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error) {
	// Build cache key
	var valueIDs *[]uuid.UUID
	if param.AttributeValueIDs != nil {
		valueIDs = param.AttributeValueIDs
	}
	searchStr := ""
	if param.Search != nil {
		searchStr = *param.Search
	}
	attributeID := uuid.Nil
	if param.AttributeID != nil {
		attributeID = *param.AttributeID
	}
	cacheKey := constant.AttributeValueListKey(attributeID, valueIDs, searchStr, *param.Limit, *param.Page)

	// Try to get from cache
	if a.redisClient != nil {
		cachedData, err := a.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var pagination Pagination[domain.AttributeValue]
			if err := json.Unmarshal([]byte(cachedData), &pagination); err == nil {
				return &pagination, nil
			}
		}
	}

	attribute, err := a.attributeRepo.ListValues(
		ctx,
		param.AttributeID,
		param.AttributeValueIDs,
		param.Search,
		*param.Limit,
		*param.Page,
	)
	if err != nil {
		return nil, err
	}
	count, err := a.attributeRepo.CountValues(
		ctx,
		param.AttributeID,
		param.AttributeValueIDs,
	)
	if err != nil {
		return nil, err
	}
	pagination := newPagination(*attribute, *count, *param.Page, *param.Limit)

	// Cache the result
	if a.redisClient != nil {
		if data, err := json.Marshal(pagination); err == nil {
			a.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLAttributeValue)*time.Second)
		}
	}

	return pagination, nil
}

func (a *AttributeImpl) Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	err = a.attributeService.Update(
		attribute,
		param.Data.Name,
	)
	if err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	if a.redisClient != nil {
		a.redisClient.Del(ctx, constant.AttributeGetKey(param.AttributeID))
		iter := a.redisClient.Scan(ctx, 0, constant.AttributeListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
	}

	return attribute, nil
}

func (a *AttributeImpl) UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*domain.AttributeValue, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	err = a.attributeService.UpdateValue(
		*attribute,
		param.AttributeValueID,
		param.Data.Value,
	)
	if err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}
	attributeValue := attribute.GetValueByID(param.AttributeValueID)
	if attributeValue == nil {
		return nil, domain.ErrNotFound
	}

	// Invalidate cache
	if a.redisClient != nil {
		a.redisClient.Del(ctx, constant.AttributeGetKey(param.AttributeID))
		iter := a.redisClient.Scan(ctx, 0, constant.AttributeListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
		iter = a.redisClient.Scan(ctx, 0, constant.AttributeValueListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
	}

	return attributeValue, nil
}

func (a *AttributeImpl) Delete(ctx context.Context, param DeleteAttributeParam) error {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return err
	}
	err = a.attributeService.Remove(attribute)
	if err != nil {
		return err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return err
	}

	// Invalidate cache
	if a.redisClient != nil {
		a.redisClient.Del(ctx, constant.AttributeGetKey(param.AttributeID))
		iter := a.redisClient.Scan(ctx, 0, constant.AttributeListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
	}

	return nil
}

func (a *AttributeImpl) DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return err
	}
	err = a.attributeService.RemoveValue(*attribute, param.AttributeValueID)
	if err != nil {
		return err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return err
	}

	// Invalidate cache
	if a.redisClient != nil {
		a.redisClient.Del(ctx, constant.AttributeGetKey(param.AttributeID))
		iter := a.redisClient.Scan(ctx, 0, constant.AttributeListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
		iter = a.redisClient.Scan(ctx, 0, constant.AttributeValueListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
	}

	return nil
}
