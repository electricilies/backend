package application

import (
	"context"

	"backend/internal/domain"
)

type AttributeImpl struct {
	attributeRepo    domain.AttributeRepository
	attributeService domain.AttributeService
}

func ProvideAttribute(attributeRepo domain.AttributeRepository) *AttributeImpl {
	return &AttributeImpl{
		attributeRepo: attributeRepo,
	}
}

var _ Attribute = &AttributeImpl{}

func (a *AttributeImpl) Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	attribute, err := a.attributeService.Create(param.Data.Name, param.Data.Code)
	if err != nil {
		return nil, err
	}
	err = a.attributeRepo.Save(ctx, attribute)
	if err != nil {
		return nil, err
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
	attribute.AddValues(*attributeValue)
	err = a.attributeRepo.Save(ctx, attribute)
	if err != nil {
		return nil, err
	}
	return attributeValue, nil
}

func (a *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error) {
	attributes, err := a.attributeRepo.List(
		ctx,
		param.IDs,
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
		param.IDs,
		param.Deleted,
	)
	if err != nil {
		return nil, err
	}
	pagination := newPagintion(*attributes, *count, *param.Page, *param.Limit)
	return pagination, nil
}

func (a *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	attribute, err := a.attributeRepo.Get(ctx, param.AttributeID)
	if err != nil {
		return nil, err
	}
	return attribute, nil
}

func (a *AttributeImpl) ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error) {
	panic("implement me")
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
	err = a.attributeRepo.Save(ctx, attribute)
	if err != nil {
		return nil, err
	}
	return attribute, nil
}

func (a *AttributeImpl) UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}

func (a *AttributeImpl) Delete(ctx context.Context, param DeleteAttributeParam) error {
	return a.attributeRepo.Remove(ctx, param.AttributeID)
}

func (a *AttributeImpl) DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error {
	// attribute , err := a.attributeRepo.Get(ctx, param.AttributeID)
	// if err != nil {
	// 	return err
	// }
	panic("implement me")
}
