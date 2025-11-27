package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type OrderResponseDto struct {
	ID          uuid.UUID              `json:"id"           binding:"required"`
	Address     string                 `json:"address"      binding:"required"`
	Provider    domain.OrderProvider   `json:"provider"     binding:"required"`
	Status      domain.OrderStatus     `json:"status"       binding:"required"`
	IsPaid      bool                   `json:"is_paid"      binding:"required"`
	CreatedAt   string                 `json:"created_at"   binding:"required"`
	UpdatedAt   string                 `json:"updated_at"   binding:"required"`
	Items       []OrderItemResponseDto `json:"items"        binding:"omitempty,dive"`
	TotalAmount int64                  `json:"total_amount" binding:"required"`
	UserID      uuid.UUID              `json:"user_id"      binding:"required"`
}

type OrderItemResponseDto struct {
	ID       uuid.UUID `json:"id"       binding:"required"`
	Quantity int       `json:"quantity" binding:"required,gt=0"`
	Price    int64     `json:"price"    binding:"required,gt=0"`
}

type OrderItemProductResponseDto struct {
	ID     uuid.UUID `json:"id"     binding:"required"`
	Name   string    `json:"name"   binding:"required"`
	Price  float64   `json:"price"  binding:"required"`
	Rating float64   `json:"rating" binding:"required"`
}

type OrderItemProductVariantResponseDto struct {
	ID  uuid.UUID `json:"id"  binding:"required"`
	SKU string    `json:"sku" binding:"required"`
}
