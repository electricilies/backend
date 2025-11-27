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
	Address     string                `json:"address"     binding:"required"`
	Provider    domain.OrderProvider  `json:"provider"    binding:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	Items       []CreateOrderItemData `json:"items"       binding:"required,dive"`
	TotalAmount int64                 `json:"totalAmount" binding:"required"`
}

type CreateOrderItemData struct {
	ProductID        uuid.UUID `json:"productId"        binding:"required"`
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required"`
	Price            int64     `json:"price"            binding:"required"`
}

type UpdateOrderRequestDto struct {
	OrderID uuid.UUID
	Data    UpdateOrderData
}

type UpdateOrderData struct {
	Status  domain.OrderStatus `json:"status,omitempty"  binding:"omitempty,oneof=Pending Processing Shipped Delivered Cancelled"`
	Address string             `json:"address,omitempty"`
}

type GetOrderRequestDto struct {
	OrderID uuid.UUID
}

type DeleteOrderRequestDto struct {
	OrderID uuid.UUID
}
