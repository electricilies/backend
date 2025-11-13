package application

import (
	"context"

	"backend/internal/domain/cart"
)

type Cart interface {
	GetCartByUser(ctx context.Context, userID string, queryParams *cart.QueryParams) (*cart.Model, error)
}

type cartApp struct {
	cartRepo cart.Repository
}

func NewCart(cartRepo cart.Repository) Cart {
	return &cartApp{
		cartRepo: cartRepo,
	}
}

func (a *cartApp) GetCartByUser(ctx context.Context, userID string, queryParams *cart.QueryParams) (*cart.Model, error) {
	return a.cartRepo.GetCartByUser(ctx, userID, queryParams)
}
