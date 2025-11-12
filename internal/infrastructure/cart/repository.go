package cart

import (
	"context"

	"backend/internal/domain/cart"
	"backend/internal/domain/param"
)

type repository struct{}

func NewRepository() cart.Repository {
	return &repository{}
}

func (r *repository) GetCartByUser(ctx context.Context, userID string, paginationParams *param.Pagination) (*cart.Model, error) {
	return &cart.Model{}, nil
}
