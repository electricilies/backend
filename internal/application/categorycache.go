package application

import (
	"context"

	"backend/internal/delivery/http"

	"github.com/google/uuid"
)

type CategoryCache interface {
	Get(ctx context.Context, param CategoryCacheParam) (*http.CategoryResponseDto, error)
	Set(ctx context.Context, param CategoryCacheParam, category *http.CategoryResponseDto) error
	Invalidate(ctx context.Context, param CategoryCacheParam) error
	GetList(ctx context.Context, param CategoryCacheListParam) (*http.PaginationResponseDto[http.CategoryResponseDto], error)
	SetList(ctx context.Context, param CategoryCacheListParam, pagination *http.PaginationResponseDto[http.CategoryResponseDto]) error
	InvalidateList(ctx context.Context, param CategoryCacheListParam) error
	InvalidateAlls(ctx context.Context) error
}

type CategoryCacheParam struct {
	ID uuid.UUID
}

type CategoryCacheListParam struct {
	Search string
	Limit  int
	Page   int
}
