package application

import (
	"backend/internal/domain"
	"context"
)

type Cart interface {
	Get(ctx context.Context, param GetCartParam) (*domain.Cart, error)
	Create(ctx context.Context, param CreateCartParam) (*Cart, error)
	CreateItem(ctx context.Context, param CreateCartItemParam) (*domain.CartItem, error)
	UpdateItem(ctx context.Context, param UpdateCartItemParam) (*domain.CartItem, error)
	DeleteItem(ctx context.Context, param DeleteCartItemParam) error
}
