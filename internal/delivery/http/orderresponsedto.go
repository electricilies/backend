package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type OrderResponseDto struct {
	ID           uuid.UUID              `json:"id"            binding:"required"`
	RecipentName string                 `json:"recipent_name" binding:"required"`
	PhoneNumber  string                 `json:"phone_number"  binding:"required"`
	Address      string                 `json:"address"       binding:"required"`
	Provider     domain.OrderProvider   `json:"provider"      binding:"required"`
	Status       domain.OrderStatus     `json:"status"        binding:"required"`
	IsPaid       bool                   `json:"is_paid"       binding:"required"`
	CreatedAt    time.Time              `json:"created_at"    binding:"required"`
	UpdatedAt    time.Time              `json:"updated_at"    binding:"required"`
	Items        []OrderItemResponseDto `json:"items"         binding:"omitempty,dive"`
	TotalAmount  int64                  `json:"total_amount"  binding:"required"`
	UserID       uuid.UUID              `json:"user_id"       binding:"required"`
}

type OrderItemResponseDto struct {
	ID             uuid.UUID                          `json:"id"             binding:"required"`
	Product        OrderItemProductResponseDto        `json:"product"        binding:"required"`
	ProductVariant OrderItemProductVariantResponseDto `json:"productVariant" binding:"required"`
	Quantity       int                                `json:"quantity"       binding:"required,gt=0"`
	Price          int64                              `json:"price"          binding:"required,gt=0"`
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

// ToOrderResponseDto maps a domain.Order to OrderResponseDto
// Note: Items will be empty initially, use WithOrderItems to populate with product data
func ToOrderResponseDto(order *domain.Order) *OrderResponseDto {
	if order == nil {
		return nil
	}

	return &OrderResponseDto{
		ID:           order.ID,
		RecipentName: order.RecipientName,
		PhoneNumber:  order.PhoneNumber,
		Address:      order.Address,
		Provider:     order.Provider,
		Status:       order.Status,
		IsPaid:       order.IsPaid,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		TotalAmount:  order.TotalAmount,
		UserID:       order.UserID,
	}
}

// WithOrderItems enriches OrderResponseDto with order items that have product and variant data
func (o *OrderResponseDto) WithOrderItems(items []OrderItemResponseDto) *OrderResponseDto {
	o.Items = items
	return o
}

// ToOrderItemResponseDto creates an OrderItemResponseDto from domain.OrderItem with product and variant
func ToOrderItemResponseDto(
	item *domain.OrderItem,
	product *domain.Product,
	variant *domain.ProductVariant,
) *OrderItemResponseDto {
	if item == nil {
		return nil
	}

	itemDto := &OrderItemResponseDto{
		ID:       item.ID,
		Quantity: item.Quantity,
		Price:    item.Price,
	}

	if product != nil {
		itemDto.Product = OrderItemProductResponseDto{
			ID:     product.ID,
			Name:   product.Name,
			Price:  float64(product.Price),
			Rating: product.Rating,
		}
	}

	if variant != nil {
		itemDto.ProductVariant = OrderItemProductVariantResponseDto{
			ID:  variant.ID,
			SKU: variant.SKU,
		}
	}

	return itemDto
}
