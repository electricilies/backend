package http

import (
	"context"

	"backend/internal/domain"
)

type CartApplication interface {
	Get(ctx context.Context, param GetCartRequestDto) (*domain.Cart, error)
	Create(ctx context.Context, param CreateCartRequestDto) (*domain.Cart, error)
	CreateItem(ctx context.Context, param CreateCartItemRequestDto) (*domain.CartItem, error)
	UpdateItem(ctx context.Context, param UpdateCartItemRequestDto) (*domain.CartItem, error)
	DeleteItem(ctx context.Context, param DeleteCartItemRequestDto) error
}
