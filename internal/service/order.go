package service

import (
	"context"

	"backend/internal/domain"
)

type Order interface {
	Create(ctx context.Context, param CreateOrderParam) (*domain.Order, error)
	Update(ctx context.Context, param UpdateOrderParam) (*domain.Order, error)
	List(ctx context.Context, param ListOrderParam) (*Pagination[domain.Order], error)
	Get(ctx context.Context, param GetOrderParam) (*domain.Order, error)
	Delete(ctx context.Context, param DeleteOrderParam) error
}
