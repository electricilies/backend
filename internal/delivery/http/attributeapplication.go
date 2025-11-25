package http

import (
	"context"

	"backend/internal/domain"
)

type AttributeApplication interface {
	Create(ctx context.Context, param CreateAttributeRequestDto) (*domain.Attribute, error)
	CreateValue(ctx context.Context, param CreateAttributeValueRequestDto) (*domain.AttributeValue, error)
	List(ctx context.Context, param ListAttributesRequestDto) (*PaginationResponseDto[domain.Attribute], error)
	Get(ctx context.Context, param GetAttributeRequestDto) (*domain.Attribute, error)
	ListValues(ctx context.Context, param ListAttributeValuesRequestDto) (*PaginationResponseDto[domain.AttributeValue], error)
	Update(ctx context.Context, param UpdateAttributeRequestDto) (*domain.Attribute, error)
	UpdateValue(ctx context.Context, param UpdateAttributeValueRequestDto) (*domain.AttributeValue, error)
	Delete(ctx context.Context, param DeleteAttributeRequestDto) error
	DeleteValue(ctx context.Context, param DeleteAttributeValueRequestDto) error
}
