package domain

import (
	"context"
)

type CartService interface {
	Get(ctx context.Context, param GetCartParam) (*Cart, error)
	Create(ctx context.Context, param CreateCartParam) (*Cart, error)
	CreateItem(ctx context.Context, param CreateCartItemParam) (*CartItem, error)
	UpdateItem(ctx context.Context, param UpdateCartItemParam) (*CartItem, error)
	DeleteItem(ctx context.Context, param DeleteCartItemParam) error
}
