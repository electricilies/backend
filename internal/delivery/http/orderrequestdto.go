package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListOrderRequestDto struct {
	PaginationRequestDto
	IDs       *[]uuid.UUID `binding:"omitnil"`
	UserIDs   *[]uuid.UUID `binding:"omitnil"`
	StatusIDs *[]uuid.UUID `binding:"omitnil"`
}

type CreateOrderRequestDto struct {
	UserID uuid.UUID       `binding:"required"`
	Data   CreateOrderData `binding:"required"`
}

type CreateOrderData struct {
	Address     string                `json:"address"     binding:"required"`
	Provider    domain.OrderProvider  `json:"provider"    binding:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	Items       []CreateOrderItemData `json:"items"       binding:"required,dive"`
	TotalAmount int64                 `json:"totalAmount" binding:"required"`
}

type CreateOrderItemData struct {
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required"`
	Price            int64     `json:"price"            binding:"required"`
}

type UpdateOrderRequestDto struct {
	OrderID uuid.UUID       `binding:"required"`
	Data    UpdateOrderData `binding:"required"`
}

type UpdateOrderData struct {
	Status  *domain.OrderStatus `json:"status"  binding:"omitnil,oneof=Pending Processing Shipped Delivered Cancelled"`
	Address *string             `json:"address" binding:"omitnil"`
}

type GetOrderRequestDto struct {
	OrderID uuid.UUID `binding:"required"`
}

type DeleteOrderRequestDto struct {
	OrderID uuid.UUID `binding:"required"`
}
