package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Order struct{}

func ProvideOrder() *Order {
	return &Order{}
}

var _ domain.OrderService = &Order{}

func (s *Order) Create(ctx context.Context, param domain.CreateOrderParam) (*domain.Order, error) {
	panic("implement me")
}

func (s *Order) Update(ctx context.Context, param domain.UpdateOrderParam) (*domain.Order, error) {
	panic("implement me")
}

func (s *Order) Get(ctx context.Context, param domain.GetOrderParam) (*domain.Order, error) {
	panic("implement me")
}

func (s *Order) Delete(ctx context.Context, param domain.DeleteOrderParam) error {
	panic("implement me")
}
