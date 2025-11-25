package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"
)

type Attribute struct {
	attributeRepo    domain.AttributeRepository
	attributeService domain.AttributeService
	attributeCache   AttributeCache
}

func ProvideAttribute(attributeRepo domain.AttributeRepository, attributeService domain.AttributeService, attributeCache AttributeCache) *Attribute {
	return &Attribute{
		attributeRepo:    attributeRepo,
		attributeService: attributeService,
		attributeCache:   attributeCache,
	}
}

var _ http.AttributeApplication = &Attribute{}

func (a *Attribute) Create(ctx context.Context, param http.CreateAttributeRequestDto) (*http.AttributeResponseDto, error) {
	attribute, err := domain.NewAttribute(param.Data.Code, param.Data.Name)
	if err != nil {
		return nil, err
	}
	if err := a.attributeService.Validate(*attribute); err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	_ = a.attributeCache.InvalidateAttributeList(ctx)

	return http.ToAttributeResponseDto(attribute), nil
}

func (a *Attribute) CreateValue(ctx context.Context, param http.CreateAttributeValueRequestDto) (*http.AttributeValueResponseDto, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	attributeValue, err := domain.NewAttributeValue(param.Data.Value)
	if err != nil {
		return nil, err
	}
	attribute.AddValues(*attributeValue)
	if err := a.attributeService.Validate(*attribute); err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	_ = a.attributeCache.InvalidateAllAttributes(ctx)

	return http.ToAttributeValueResponseDto(attributeValue), nil
}

func (a *Attribute) List(ctx context.Context, param http.ListAttributesRequestDto) (*http.PaginationResponseDto[http.AttributeResponseDto], error) {
	cacheKey := a.attributeCache.BuildListCacheKey(
		param.AttributeIDs,
		param.Search,
		param.Deleted,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if cachedPagination, err := a.attributeCache.GetAttributeList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}

	attributes, err := a.attributeRepo.List(
		ctx,
		param.AttributeIDs,
		nil,
		param.Search,
		param.Deleted,
		param.Limit,
		(param.Page-1)*param.Limit,
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

	// Map domain models to response DTOs
	attributeDtos := http.ToAttributeResponseDtoList(*attributes)

	pagination := newPaginationResponseDto(
		attributeDtos,
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

func (a *Attribute) Get(ctx context.Context, param http.GetAttributeRequestDto) (*http.AttributeResponseDto, error) {
	if cachedAttribute, err := a.attributeCache.GetAttribute(ctx, param.AttributeID); err == nil {
		return cachedAttribute, nil
	}

	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}

	attributeDto := http.ToAttributeResponseDto(attribute)
	_ = a.attributeCache.SetAttribute(ctx, param.AttributeID, attributeDto)

	return attributeDto, nil
}

func (a *Attribute) ListValues(ctx context.Context, param http.ListAttributeValuesRequestDto) (*http.PaginationResponseDto[http.AttributeValueResponseDto], error) {
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
	attributeValues, err := a.attributeRepo.ListValues(
		ctx,
		param.AttributeID,
		param.AttributeValueIDs,
		param.Search,
		param.Deleted,
		param.Limit,
		(param.Page-1)*param.Limit,
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

	// Map domain models to response DTOs
	attributeValueDtos := http.ToAttributeValueResponseDtoList(*attributeValues)

	pagination := newPaginationResponseDto(
		attributeValueDtos,
		*count,
		param.Page,
		param.Limit,
	)
	_ = a.attributeCache.SetAttributeValueList(ctx, cacheKey, pagination)
	return pagination, nil
}

func (a *Attribute) Update(ctx context.Context, param http.UpdateAttributeRequestDto) (*http.AttributeResponseDto, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	attribute.Update(param.Data.Name)
	if err := a.attributeService.Validate(*attribute); err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return nil, err
	}
	_ = a.attributeCache.InvalidateAttribute(ctx, param.AttributeID)
	_ = a.attributeCache.InvalidateAttributeList(ctx)
	return http.ToAttributeResponseDto(attribute), nil
}

func (a *Attribute) UpdateValue(ctx context.Context, param http.UpdateAttributeValueRequestDto) (*http.AttributeValueResponseDto, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	if err := attribute.UpdateValue(param.AttributeValueID, param.Data.Value); err != nil {
		return nil, err
	}
	if err := a.attributeService.Validate(*attribute); err != nil {
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
	return http.ToAttributeValueResponseDto(attributeValue), nil
}

func (a *Attribute) Delete(ctx context.Context, param http.DeleteAttributeRequestDto) error {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return err
	}
	attribute.Remove()
	if err := a.attributeService.Validate(*attribute); err != nil {
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

func (a *Attribute) DeleteValue(ctx context.Context, param http.DeleteAttributeValueRequestDto) error {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return err
	}
	if err := attribute.RemoveValue(param.AttributeValueID); err != nil {
		return err
	}
	if err := a.attributeService.Validate(*attribute); err != nil {
		return err
	}
	err = a.attributeRepo.Save(ctx, *attribute)
	if err != nil {
		return err
	}
	_ = a.attributeCache.InvalidateAllAttributes(ctx)
	return nil
}
