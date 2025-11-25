package http

import (
	"context"
)

type CategoryApplication interface {
	Create(ctx context.Context, param CreateCategoryRequestDto) (*CategoryResponseDto, error)
	List(ctx context.Context, param ListCategoryRequestDto) (*PaginationResponseDto[CategoryResponseDto], error)
	Get(ctx context.Context, param GetCategoryRequestDto) (*CategoryResponseDto, error)
	Update(ctx context.Context, param UpdateCategoryRequestDto) (*CategoryResponseDto, error)
}
