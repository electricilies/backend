package service

import (
	"backend/internal/domain"
)

type Order struct{}

func ProvideOrder() *Order {
	return &Order{}
}

var _ domain.OrderService = &Order{}
