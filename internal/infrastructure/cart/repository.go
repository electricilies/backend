package cart

import (
	"backend/internal/domain/cart"
	"backend/internal/domain/pagination"
)

type repository struct{}

func NewRepository() cart.Repository {
	return &repository{}
}

func (r *repository) GetCartByUser(userID string, paginationParams *pagination.Params) (*cart.Model, error) {
	return &cart.Model{}, nil
}
