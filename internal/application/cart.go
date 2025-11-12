package application

import (
	"context"

	"backend/internal/domain/cart"
	"backend/internal/domain/param"
)

type Cart interface {
	GetCartByUser(ctx context.Context, userID string, paginationParams *param.Pagination) (*cart.Model, error)
}

type cartApp struct {
	cartRepo cart.Repository
}

func NewCart(cartRepo cart.Repository) Cart {
	return &cartApp{
		cartRepo: cartRepo,
	}
}

func (a *cartApp) GetCartByUser(ctx context.Context, userID string, paginationParams *param.Pagination) (*cart.Model, error) {
	return a.cartRepo.GetCartByUser(ctx, userID, paginationParams)
}
