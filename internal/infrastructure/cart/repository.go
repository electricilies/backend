package cart

import (
	"context"

	"backend/internal/domain/cart"
)

type RepositoryImpl struct{}

func NewRepository() cart.Repository {
	return &RepositoryImpl{}
}

func ProvideRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error) {
	return &cart.Model{}, nil
}

func (r *RepositoryImpl) AddItem(ctx context.Context, itemModel *cart.ItemModel) (*cart.ItemModel, error) {
	return &cart.ItemModel{}, nil
}

func (r *RepositoryImpl) UpdateItem(
	ctx context.Context,
	itemModel *cart.ItemModel,
	id string,
) (*cart.ItemModel, error) {
	return &cart.ItemModel{}, nil
}

func (r *RepositoryImpl) RemoveItem(ctx context.Context, id string) error {
	return nil
}
