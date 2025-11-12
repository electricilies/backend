package cart

import (
	"context"

	"backend/internal/domain/cart"
	"backend/internal/domain/params"
)

type repository struct{}

func NewRepository() cart.Repository {
	return &repository{}
}

func (r *repository) GetCartByUser(ctx context.Context, userID string, paginationParams *params.Params) (*cart.Model, error) {
	return &cart.Model{}, nil
}
