package application

import (
	"context"

	"backend/internal/domain"
)

type AttributeImpl struct{}

func ProvideAttribute() *AttributeImpl {
	return &AttributeImpl{}
}

var _ Attribute = &AttributeImpl{}

func (a *AttributeImpl) Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (a *AttributeImpl) CreateValue(ctx context.Context, param CreateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}

func (a *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error) {
	panic("implement me")
}

func (a *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (a *AttributeImpl) ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error) {
	panic("implement me")
}

func (a *AttributeImpl) Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (a *AttributeImpl) UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}

func (a *AttributeImpl) Delete(ctx context.Context, param DeleteAttributeParam) error {
	panic("implement me")
}

func (a *AttributeImpl) DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error {
	panic("implement me")
}
