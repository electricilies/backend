package service

import (
	"context"

	"backend/internal/domain"
)

type Attribute interface {
	Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error)
	List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error)
	Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error)
	Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error)
	Delete(ctx context.Context, param DeleteAttributeParam) error
	CreateValue(ctx context.Context, param CreateAttributeValueParam) (*domain.AttributeValue, error)
	UpdateValues(ctx context.Context, param UpdateAttributeValuesParam) (*[]domain.AttributeValue, error)
	ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error)
	DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error
}
