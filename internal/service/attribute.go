package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Attribute struct{}

func ProvideAttribute() *Attribute {
	return &Attribute{}
}

var _ domain.AttributeService = &Attribute{}

func (s *Attribute) Get(ctx context.Context, param domain.GetAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (s *Attribute) Create(ctx context.Context, param domain.CreateAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (s *Attribute) Update(ctx context.Context, param domain.UpdateAttributeParam) (*domain.Attribute, error) {
	panic("implement me")
}

func (s *Attribute) Delete(ctx context.Context, param domain.DeleteAttributeParam) error {
	panic("implement me")
}

func (s *Attribute) CreateValue(ctx context.Context, param domain.CreateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}

func (s *Attribute) UpdateValue(ctx context.Context, param domain.UpdateAttributeValueParam) (*domain.AttributeValue, error) {
	panic("implement me")
}

func (s *Attribute) DeleteValue(ctx context.Context, param domain.DeleteAttributeValueParam) error {
	panic("implement me")
}
