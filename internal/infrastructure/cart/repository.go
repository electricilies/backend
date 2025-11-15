package cart

import (
	"context"

	"backend/internal/domain/cart"
)

type repositoryImpl struct{}

func NewRepository() cart.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error) {
	return &cart.Model{}, nil
}

func (r *repositoryImpl) AddItem(ctx context.Context, cartItemModel *cart.ItemModel) (*cart.ItemModel, error) {
	return &cart.ItemModel{}, nil
}

func (r *repositoryImpl) UpdateItem(ctx context.Context, cartItemModel *cart.ItemModel, id string) (*cart.ItemModel, error) {
	return &cart.ItemModel{}, nil
}

func (r *repositoryImpl) RemoveItem(ctx context.Context, id string) error {
	return nil
}
