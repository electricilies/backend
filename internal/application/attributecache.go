package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type AttributeCache interface {
	Get(ctx context.Context, param AttributeCacheParam) (*http.AttributeResponseDto, error)
	Set(ctx context.Context, param AttributeCacheParam, attribute *http.AttributeResponseDto) error
	Invalidate(ctx context.Context, param AttributeCacheParam) error
	GetList(ctx context.Context, param AttributeCacheListParam) (*http.PaginationResponseDto[http.AttributeResponseDto], error)
	SetList(ctx context.Context, param AttributeCacheListParam, pagination *http.PaginationResponseDto[http.AttributeResponseDto]) error
	InvalidateList(ctx context.Context, param AttributeCacheListParam) error
	GetValueList(ctx context.Context, param AttributeCacheValueListParam) (*http.PaginationResponseDto[http.AttributeValueResponseDto], error)
	SetValueList(ctx context.Context, param AttributeCacheValueListParam, pagination *http.PaginationResponseDto[http.AttributeValueResponseDto]) error
	InvalidateValueList(ctx context.Context, param AttributeCacheValueListParam) error
	InvalidateAlls(ctx context.Context) error
}

type AttributeCacheParam struct {
	ID uuid.UUID
}

type AttributeCacheListParam struct {
	IDs     []uuid.UUID
	Search  string
	Deleted domain.DeletedParam
	Limit   int
	Page    int
}

type AttributeCacheValueListParam struct {
	ID       uuid.UUID
	ValueIDs []uuid.UUID
	Search   string
	Limit    int
	Page     int
}
