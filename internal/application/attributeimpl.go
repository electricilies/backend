package application

import (
	"context"
	"log"

	"backend/internal/domain"
)

type AttributeImpl struct {
	attributeRepo    domain.AttributeRepository
	attributeService domain.AttributeService
	attributeCache   AttributeCache
}

func ProvideAttribute(attributeRepo domain.AttributeRepository, attributeService domain.AttributeService, attributeCache AttributeCache) *AttributeImpl {
	return &AttributeImpl{
		attributeRepo:    attributeRepo,
		attributeService: attributeService,
		attributeCache:   attributeCache,
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

	_ = a.attributeCache.InvalidateAttributeList(ctx)

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
	err = a.attributeService.AddValues(attribute, *attributeValue)
	if err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	_ = a.attributeCache.InvalidateAllAttributes(ctx)

	return attributeValue, nil
}

func (a *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error) {
	cacheKey := a.attributeCache.BuildListCacheKey(
		param.AttributeIDs,
		param.Search,
		param.Deleted,
		param.Limit,
		param.Page,
	)
	log.Println("Attribute List Cache Key:", cacheKey)

	// Try to get from cache
	if cachedPagination, err := a.attributeCache.GetAttributeList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}

	attributes, err := a.attributeRepo.List(
		ctx,
		param.AttributeIDs,
		param.Search,
		param.Deleted,
		param.Limit,
		param.Page,
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
	pagination := newPagination(
		*attributes,
		*count,
		param.Page,
		param.Limit,
	)

	// Cache the result
	err = a.attributeCache.SetAttributeList(ctx, cacheKey, pagination)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

func (a *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	if cachedAttribute, err := a.attributeCache.GetAttribute(ctx, param.AttributeID); err == nil {
		return cachedAttribute, nil
	}

	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}

	_ = a.attributeCache.SetAttribute(ctx, param.AttributeID, attribute)

	return attribute, nil
}

func (a *AttributeImpl) ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error) {
	cacheKey := a.attributeCache.BuildValueListCacheKey(
		param.AttributeID,
		param.AttributeValueIDs,
		param.Search,
		param.Limit,
		param.Page,
	)
	if cachedPagination, err := a.attributeCache.GetAttributeValueList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}
	attribute, err := a.attributeRepo.ListValues(
		ctx,
		param.AttributeID,
		param.AttributeValueIDs,
		param.Search,
		param.Deleted,
		param.Limit,
		param.Page,
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
	pagination := newPagination(
		*attribute,
		*count,
		param.Page,
		param.Limit,
	)
	_ = a.attributeCache.SetAttributeValueList(ctx, cacheKey, pagination)
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
	_ = a.attributeCache.InvalidateAttribute(ctx, param.AttributeID)
	_ = a.attributeCache.InvalidateAttributeList(ctx)
	return attribute, nil
}

func (a *AttributeImpl) UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*domain.AttributeValue, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	err = a.attributeService.UpdateValue(
		attribute,
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
	_ = a.attributeCache.InvalidateAllAttributes(ctx)
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
	_ = a.attributeCache.InvalidateAttribute(ctx, param.AttributeID)
	_ = a.attributeCache.InvalidateAttributeList(ctx)
	return nil
}

func (a *AttributeImpl) DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return err
	}
	err = a.attributeService.RemoveValue(attribute, param.AttributeValueID)
	if err != nil {
		return err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return err
	}
	_ = a.attributeCache.InvalidateAllAttributes(ctx)
	return nil
}
