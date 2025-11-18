package service

import (
	"context"

	"backend/internal/domain"
)

type OrderImpl struct{}

func ProvideOrder() *OrderImpl {
	return &OrderImpl{}
}

var _ Order = &OrderImpl{}

func (o *OrderImpl) Create(ctx context.Context, param CreateOrderParam) (*domain.Order, error) {
	panic("implement me")
}

func (o *OrderImpl) Update(ctx context.Context, param UpdateOrderParam) (*domain.Order, error) {
	panic("implement me")
}

func (o *OrderImpl) List(ctx context.Context, param ListOrderParam) (*Pagination[domain.Order], error) {
	panic("implement me")
}

func (o *OrderImpl) Get(ctx context.Context, param GetOrderParam) (*domain.Order, error) {
	panic("implement me")
}

func (o *OrderImpl) Delete(ctx context.Context, param DeleteOrderParam) error {
	panic("implement me")
}
