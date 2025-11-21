package application

import (
	"context"

	"backend/internal/domain"
)

type Attribute interface {
	Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error)
	CreateValue(ctx context.Context, param CreateAttributeValueParam) (*domain.AttributeValue, error)
	List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error)
	Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error)
	ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error)
	Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error)
	UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*domain.AttributeValue, error)
	Delete(ctx context.Context, param DeleteAttributeParam) error
	DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error
}
