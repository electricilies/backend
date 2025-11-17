package service

import (
	"context"

	"backend/internal/domain"
)

type GetCartParam struct {
	CartID int `json:"cartId" binding:"required"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type AddCartItemParam struct {
	CartID    int `json:"cartId" binding:"required"`
	ProductID int `json:"productId" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type UpdateCartItemParam struct {
	CartID   int `json:"cartId" binding:"required"`
	ItemID   int `json:"itemId" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type RemoveCartItemParam struct {
	CartID int `json:"cartId" binding:"required"`
	ItemID int `json:"itemId" binding:"required"`
}

type Cart interface {
	GetCart(context.Context, GetCartParam) (*domain.Cart, error)
	AddCartItem(context.Context, AddCartItemParam) (*domain.CartItem, error)
	UpdateCartItem(context.Context, UpdateCartItemParam) (*domain.CartItem, error)
	RemoveCartItem(context.Context, RemoveCartItemParam) error
}

type CartImpl struct{}

func ProvideCart() *CartImpl {
	return &CartImpl{}
}

var _ Cart = &CartImpl{}

func (s *CartImpl) GetCart(ctx context.Context, param GetCartParam) (*domain.Cart, error) {
	return nil, nil
}

func (s *CartImpl) AddCartItem(ctx context.Context, param AddCartItemParam) (*domain.CartItem, error) {
	return nil, nil
}

func (s *CartImpl) UpdateCartItem(ctx context.Context, param UpdateCartItemParam) (*domain.CartItem, error) {
	return nil, nil
}

func (s *CartImpl) RemoveCartItem(ctx context.Context, param RemoveCartItemParam) error {
	return nil
}
