package cart

import (
	"context"

	"backend/internal/domain/cart"
)

type repository struct{}

func NewRepository() cart.Repository {
	return &repository{}
}

func (r *repository) GetCartByUser(ctx context.Context, userID string, queryParams *cart.QueryParams) (*cart.Model, error) {
	return &cart.Model{}, nil
}
