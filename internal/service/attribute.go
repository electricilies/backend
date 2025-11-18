package service

import (
	"context"

	"backend/internal/domain"
)

type Attribute interface {
	Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error)
	List(ctx context.Context, param ListAttributesParam) (*domain.Pagination[domain.Attribute], error)
	Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error)
	Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error)
	Delete(ctx context.Context, param DeleteAttributeParam) error
	CreateValues(ctx context.Context, param []CreateAttributeValueParam) error
	UpdateValues(ctx context.Context, param []UpdateAttributeValueParam) error
}

type AttributeImpl struct{}

func ProvideAttribute() *AttributeImpl {
	return &AttributeImpl{}
}

var _ Attribute = &AttributeImpl{}

func (s *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) List(
	ctx context.Context,
	param ListAttributesParam,
) (*domain.Pagination[domain.Attribute], error) {
	return nil, nil
}

func (s *AttributeImpl) Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) Delete(ctx context.Context, param DeleteAttributeParam) error {
	return nil
}

func (s *AttributeImpl) UpdateValues(ctx context.Context, param []UpdateAttributeValueParam) error {
	return nil
}
