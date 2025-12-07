package http

import (
	"context"
)

type OrderApplication interface {
	List(ctx context.Context, param ListOrderRequestDto) (*PaginationResponseDto[OrderResponseDto], error)
	Create(ctx context.Context, param CreateOrderRequestDto) (*OrderResponseDto, error)
	Get(ctx context.Context, param GetOrderRequestDto) (*OrderResponseDto, error)
	Update(ctx context.Context, param UpdateOrderRequestDto) (*OrderResponseDto, error)
	VerifyVNPayIPN(ctx context.Context, param VerifyVNPayIPNRequestDTO) (*VerifyVNPayIPNResponseDTO, error)
}
