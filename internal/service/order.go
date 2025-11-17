package service

import (
	"context"

	"backend/internal/domain"
)

type CreateOrderParam struct {
	UserID  string                 `json:"userId" binding:"required"`
	Address string                 `json:"address" binding:"required"`
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
	PaginationParam
	OrderIDs  []int
	UserIDs   []string
	StatusIDs []int
}

type GetOrderParam struct {
	OrderID int `json:"orderId" binding:"required"`
}

type DeleteOrderParam struct {
	OrderID int `json:"orderId" binding:"required"`
}

type Order interface {
	Create(context.Context, CreateOrderParam) (*domain.Order, error)
	UpdateStatus(context.Context, UpdateOrderParam) (*domain.Order, error)
	List(context.Context, ListOrderParam) (*domain.Pagination[domain.Order], error)
	Get(context.Context, GetOrderParam) (*domain.Order, error)
	Delete(context.Context, DeleteOrderParam) error
}

type OrderImpl struct{}

func ProvideOrder() *OrderImpl {
	return &OrderImpl{}
}

var _ Order = &OrderImpl{}

func (s *OrderImpl) Create(ctx context.Context, param CreateOrderParam) (*domain.Order, error) {
	return nil, nil
}

func (s *OrderImpl) UpdateStatus(ctx context.Context, param UpdateOrderParam) (*domain.Order, error) {
	return nil, nil
}

func (s *OrderImpl) List(ctx context.Context, param ListOrderParam) (*domain.Pagination[domain.Order], error) {
	return nil, nil
}

func (s *OrderImpl) Get(ctx context.Context, param GetOrderParam) (*domain.Order, error) {
	return nil, nil
}

func (s *OrderImpl) Delete(ctx context.Context, param DeleteOrderParam) error {
	return nil
}
