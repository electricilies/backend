package http

import (
	"context"
)

type CartApplication interface {
	Get(ctx context.Context, param GetCartRequestDto) (*CartResponseDto, error)
	GetByUser(ctx context.Context, param GetCartByUserRequestDto) (*CartResponseDto, error)
	Create(ctx context.Context, param CreateCartRequestDto) (*CartResponseDto, error)
	CreateItem(ctx context.Context, param CreateCartItemRequestDto) (*CartItemResponseDto, error)
	UpdateItem(ctx context.Context, param UpdateCartItemRequestDto) (*CartItemResponseDto, error)
	DeleteItem(ctx context.Context, param DeleteCartItemRequestDto) error
}
