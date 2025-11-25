package http

import (
	"context"

	"backend/internal/domain"
)

type OrderApplication interface {
	Create(ctx context.Context, param CreateOrderRequestDto) (*domain.Order, error)
	Update(ctx context.Context, param UpdateOrderRequestDto) (*domain.Order, error)
	Get(ctx context.Context, param GetOrderRequestDto) (*domain.Order, error)
	Delete(ctx context.Context, param DeleteOrderRequestDto) error
	List(ctx context.Context, param ListOrderRequestDto) (*PaginationResponseDto[domain.Order], error)
}
