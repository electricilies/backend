package http

import (
	"context"

	"backend/internal/domain"
)

type OrderApplication interface {
	List(ctx context.Context, param ListOrderRequestDto) (*PaginationResponseDto[domain.Order], error)
	Create(ctx context.Context, param CreateOrderRequestDto) (*OrderResponseDto, error)
	Get(ctx context.Context, param GetOrderRequestDto) (*OrderResponseDto, error)
	Update(ctx context.Context, param UpdateOrderRequestDto) (*OrderResponseDto, error)
}
