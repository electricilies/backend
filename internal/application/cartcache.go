package application

import (
	"context"

	"backend/internal/delivery/http"

	"github.com/google/uuid"
)

type CartCache interface {
	Get(ctx context.Context, param CartCacheParam) (*http.CartResponseDto, error)
	Set(ctx context.Context, param CartCacheParam, cart *http.CartResponseDto) error
	Invalidate(ctx context.Context, param CartCacheParam) error
	InvalidateAlls(ctx context.Context) error
}

type CartCacheParam struct {
	ID uuid.UUID
}
