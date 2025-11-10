package application

import (
	"backend/internal/domain/cart"
	"backend/internal/domain/pagination"
)

type Cart interface {
	GetCartByUser(userID string, paginationParams *pagination.Params) (*cart.Model, error)
}

type cartApp struct {
	cartRepo cart.Repository
}

func NewCart(cartRepo cart.Repository) Cart {
	return &cartApp{
		cartRepo: cartRepo,
	}
}

func (a *cartApp) GetCartByUser(userID string, paginationParams *pagination.Params) (*cart.Model, error) {
	return a.cartRepo.GetCartByUser(userID, paginationParams)
}
