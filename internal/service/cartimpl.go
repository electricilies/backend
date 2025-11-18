package service

import (
	"context"

	"backend/internal/domain"
)

type CartImpl struct{}

func ProvideCart() *CartImpl {
	return &CartImpl{}
}

var _ Cart = &CartImpl{}

func (s *CartImpl) Get(ctx context.Context, param GetCartParam) (*domain.Cart, error) {
	panic("implement me")
}

func (s *CartImpl) Create(ctx context.Context, param CreateCartParam) (*domain.Cart, error) {
	panic("implement me")
}

func (s *CartImpl) CreateItem(ctx context.Context, param CreateCartItemParam) (*domain.CartItem, error) {
	panic("implement me")
}

func (s *CartImpl) UpdateItem(ctx context.Context, param UpdateCartItemParam) (*domain.CartItem, error) {
	panic("implement me")
}

func (s *CartImpl) DeleteItem(ctx context.Context, param DeleteCartItemParam) error {
	panic("implement me")
}
