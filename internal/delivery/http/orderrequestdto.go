package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListOrderRequestDto struct {
	PaginationRequestDto
	IDs       []uuid.UUID
	UserIDs   []uuid.UUID
	StatusIDs []uuid.UUID
}

type CreateOrderRequestDto struct {
	UserID uuid.UUID
	Data   CreateOrderData
}

type CreateOrderData struct {
	Address  string                `json:"address"  binding:"required"`
	Provider domain.OrderProvider  `json:"provider" binding:"required"`
	Items    []CreateOrderItemData `json:"items"    binding:"required,dive"`
}

type CreateOrderItemData struct {
	ProductID        uuid.UUID `json:"productId"        binding:"required"`
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required"`
}

type UpdateOrderRequestDto struct {
	OrderID uuid.UUID
	Data    UpdateOrderData
}

type UpdateOrderData struct {
	Address string             `json:"address" binding:"required"`
	Status  domain.OrderStatus `json:"status"  binding:"required"`
	IsPaid  bool               `json:"is_paid" binding:"required"`
}

type GetOrderRequestDto struct {
	OrderID uuid.UUID
}

type DeleteOrderRequestDto struct {
	OrderID uuid.UUID
}
