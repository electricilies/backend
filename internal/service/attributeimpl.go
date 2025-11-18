package service

import (
	"context"

	"backend/internal/domain"
)

type AttributeImpl struct{}

func ProvideAttribute() *AttributeImpl {
	return &AttributeImpl{}
}

var _ Attribute = &AttributeImpl{}

func (s *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (s *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error) {
	panic("implement me")
}

func (s *AttributeImpl) Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (s *AttributeImpl) Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (s *AttributeImpl) Delete(ctx context.Context, param DeleteAttributeParam) error {
	panic("implement me")
}

func (s *AttributeImpl) CreateValue(ctx context.Context, param CreateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}

func (s *AttributeImpl) UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}
