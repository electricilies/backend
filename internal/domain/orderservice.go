package domain

import (
	"context"
)

type OrderService interface {
	Create(ctx context.Context, param CreateOrderParam) (*Order, error)
	Update(ctx context.Context, param UpdateOrderParam) (*Order, error)
	Get(ctx context.Context, param GetOrderParam) (*Order, error)
	Delete(ctx context.Context, param DeleteOrderParam) error
}
