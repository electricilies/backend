package order

import (
	"context"

	"backend/internal/domain/order"
)

type repositoryImpl struct{}

func NewRepository() order.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) List(ctx context.Context) (*[]order.Model, error) {
	return &[]order.Model{}, nil
}

func (r *repositoryImpl) Create(ctx context.Context, orderModel *order.Model) (*order.Model, error) {
	return orderModel, nil
}

func (r *repositoryImpl) Update(ctx context.Context, orderModel *order.Model, id int) (*order.Model, error) {
	return orderModel, nil
}

func (r *repositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}
