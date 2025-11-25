package http

import (
	"context"

	"backend/internal/domain"
)

type CategoryApplication interface {
	Create(ctx context.Context, param CreateCategoryRequestDto) (*domain.Category, error)
	List(ctx context.Context, param ListCategoryRequestDto) (*PaginationResponseDto[domain.Category], error)
	Get(ctx context.Context, param GetCategoryRequestDto) (*domain.Category, error)
	Update(ctx context.Context, param UpdateCategoryRequestDto) (*domain.Category, error)
}
