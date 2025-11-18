package service

import (
	"context"

	"backend/internal/domain"
)

type Cart interface {
	Get(ctx context.Context, param GetCartParam) (*domain.Cart, error)
	AddItem(ctx context.Context, param AddCartItemParam) (*domain.CartItem, error)
	UpdateItem(ctx context.Context, param UpdateCartItemParam) (*domain.CartItem, error)
	RemoveItem(ctx context.Context, param RemoveCartItemParam) error
}
