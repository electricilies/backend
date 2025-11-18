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
	return nil, nil
}

func (s *CartImpl) AddItem(ctx context.Context, param AddCartItemParam) (*domain.CartItem, error) {
	return nil, nil
}

func (s *CartImpl) UpdateItem(ctx context.Context, param UpdateCartItemParam) (*domain.CartItem, error) {
	return nil, nil
}

func (s *CartImpl) RemoveItem(ctx context.Context, param RemoveCartItemParam) error {
	return nil
}
