package application

import (
	"context"

	"backend/internal/domain/cart"
)

type Cart interface {
	Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error)
}

type cartApp struct {
	cartRepo cart.Repository
}

func NewCart(cartRepo cart.Repository) Cart {
	return &cartApp{
		cartRepo: cartRepo,
	}
}

func (a *cartApp) Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error) {
	return a.cartRepo.Get(ctx, id, queryParams)
}
