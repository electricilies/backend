package application

import (
	"context"

	"backend/internal/domain/cart"
)

type Cart interface {
	Get(context.Context, int, *cart.QueryParams) (*cart.Model, error)
}

type CartImpl struct {
	cartRepo cart.Repository
}

func NewCart(cartRepo cart.Repository) Cart {
	return &CartImpl{
		cartRepo: cartRepo,
	}
}

func ProvideCart(
	cartRepo cart.Repository,
) *CartImpl {
	return &CartImpl{
		cartRepo: cartRepo,
	}
}

func (a *CartImpl) Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error) {
	return a.cartRepo.Get(ctx, id, queryParams)
}
