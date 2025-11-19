package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Cart struct{}

func ProvideCart() *Cart {
	return &Cart{}
}

var _ domain.CartService = &Cart{}

func (s *Cart) Get(ctx context.Context, param domain.GetCartParam) (*domain.Cart, error) {
	panic("implement me")
}

func (s *Cart) Create(ctx context.Context, param domain.CreateCartParam) (*domain.Cart, error) {
	panic("implement me")
}

func (s *Cart) CreateItem(ctx context.Context, param domain.CreateCartItemParam) (*domain.CartItem, error) {
	panic("implement me")
}

func (s *Cart) UpdateItem(ctx context.Context, param domain.UpdateCartItemParam) (*domain.CartItem, error) {
	panic("implement me")
}

func (s *Cart) DeleteItem(ctx context.Context, param domain.DeleteCartItemParam) error {
	panic("implement me")
}
