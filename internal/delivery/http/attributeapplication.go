package http

import (
	"context"
)

type AttributeApplication interface {
	Create(ctx context.Context, param CreateAttributeRequestDto) (*AttributeResponseDto, error)
	CreateValue(ctx context.Context, param CreateAttributeValueRequestDto) (*AttributeValueResponseDto, error)
	List(ctx context.Context, param ListAttributesRequestDto) (*PaginationResponseDto[AttributeResponseDto], error)
	Get(ctx context.Context, param GetAttributeRequestDto) (*AttributeResponseDto, error)
	ListValues(ctx context.Context, param ListAttributeValuesRequestDto) (*PaginationResponseDto[AttributeValueResponseDto], error)
	Update(ctx context.Context, param UpdateAttributeRequestDto) (*AttributeResponseDto, error)
	UpdateValue(ctx context.Context, param UpdateAttributeValueRequestDto) (*AttributeValueResponseDto, error)
	Delete(ctx context.Context, param DeleteAttributeRequestDto) error
	DeleteValue(ctx context.Context, param DeleteAttributeValueRequestDto) error
}
