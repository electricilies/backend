package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ProductCache interface {
	Get(ctx context.Context, param ProductCacheParam) (*http.ProductResponseDto, error)
	Set(ctx context.Context, param ProductCacheParam, product *http.ProductResponseDto) error
	Invalidate(ctx context.Context, param ProductCacheParam) error
	GetList(ctx context.Context, param ProductCacheListParam) (*http.PaginationResponseDto[http.ProductResponseDto], error)
	SetList(ctx context.Context, param ProductCacheListParam, pagination *http.PaginationResponseDto[http.ProductResponseDto]) error
	InvalidateList(ctx context.Context, param ProductCacheListParam) error
	InvalidateAlls(ctx context.Context) error
}

type ProductCacheParam struct {
	ID uuid.UUID
}

type ProductCacheListParam struct {
	IDs         []uuid.UUID
	Search      string
	MinPrice    int64
	MaxPrice    int64
	Rating      float64
	CategoryIDs []uuid.UUID
	Deleted     domain.DeletedParam
	SortRating  string
	SortPrice   string
	Limit       int
	Page        int
}
