package application

import (
	"context"

	"backend/internal/domain/cart"
)

type Cart interface {
	Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error)
}

type CartApp struct {
	cartRepo cart.Repository
}

func NewCart(cartRepo cart.Repository) Cart {
	return &CartApp{
		cartRepo: cartRepo,
	}
}

func ProvideCart(
	cartRepo cart.Repository,
) *CartApp {
	return &CartApp{
		cartRepo: cartRepo,
	}
}

func (a *CartApp) Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error) {
	return a.cartRepo.Get(ctx, id, queryParams)
}
