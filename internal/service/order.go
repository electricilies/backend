package service

import (
	"context"

	"backend/internal/domain"
)

type CreateOrderParam struct {
	UserID  string                   `json:"userId" binding:"required"`
	Address string                   `json:"address" binding:"required"`
	Items   []CreateOrderItemParam `json:"items" binding:"required"`
}

type CreateOrderItemParam struct {
	ProductID        int   `json:"productId" binding:"required"`
	ProductVariantID int   `json:"productVariantId,omitempty"`
	Quantity         int   `json:"quantity" binding:"required"`
	Price            int64 `json:"price" binding:"required"`
}

type UpdateOrderParam struct {
	OrderStatus string `json:"orderStatus" binding:"required"`
}

type ListOrderParam struct {
	OrderIDs  []int
	UserIDs   []string
	StatusIDs []int
	Limit     int
	Offset    int
}

type GetOrderParam struct {
	OrderID int `json:"orderId" binding:"required"`
}

type DeleteOrderParam struct {
	OrderID int `json:"orderId" binding:"required"`
}

type Order interface {
	CreateOrder(context.Context, CreateOrderParam) (*domain.Order, error)
	UpdateOrderStatus(context.Context, UpdateOrderParam) (*domain.Order, error)
	ListOrders(context.Context, ListOrderParam) (*domain.DataPagination, error)
	GetOrder(context.Context, GetOrderParam) (*domain.Order, error)
	DeleteOrder(context.Context, DeleteOrderParam) error
}

type OrderImpl struct{}

func ProvideOrder() *OrderImpl {
	return &OrderImpl{}
}

var _ Order = &OrderImpl{}

func (s *OrderImpl) CreateOrder(ctx context.Context, param CreateOrderParam) (*domain.Order, error) {
	return nil, nil
}

func (s *OrderImpl) UpdateOrderStatus(ctx context.Context, param UpdateOrderParam) (*domain.Order, error) {
	return nil, nil
}

func (s *OrderImpl) ListOrders(ctx context.Context, param ListOrderParam) (*domain.DataPagination, error) {
	return nil, nil
}

func (s *OrderImpl) GetOrder(ctx context.Context, param GetOrderParam) (*domain.Order, error) {
	return nil, nil
}

func (s *OrderImpl) DeleteOrder(ctx context.Context, param DeleteOrderParam) error {
	return nil
}

