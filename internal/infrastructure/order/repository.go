package order

import (
	"context"

	"backend/internal/domain/order"
)

type RepositoryImpl struct{}

func NewRepository() order.Repository {
	return &RepositoryImpl{}
}

func ProvideRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) List(ctx context.Context) (*[]order.Model, error) {
	return &[]order.Model{}, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, orderModel *order.Model) (*order.Model, error) {
	return orderModel, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, orderModel *order.Model, id int) (*order.Model, error) {
	return orderModel, nil
}

func (r *RepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}
