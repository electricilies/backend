package cart

import (
	"context"

	"backend/internal/domain/cart"
)

type repositoryImpl struct{}

func NewRepository() cart.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetCartByUser(ctx context.Context, userID string, queryParams *cart.QueryParams) (*cart.Model, error) {
	return &cart.Model{}, nil
}
