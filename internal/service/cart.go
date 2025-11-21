package service

import (
	"backend/internal/domain"
)

type Cart struct{}

func ProvideCart() *Cart {
	return &Cart{}
}

var _ domain.CartService = &Cart{}
