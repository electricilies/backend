package order

import (
	"time"

	"backend/internal/domain/product"
)

type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
	Cancelled  OrderStatus = "Cancelled"
)

type Model struct {
	ID        *int
	UserID    *string
	Address   *string
	Status    *OrderStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Items     []ItemModel
	// TODO: implement payment later
	//	Payment   Payment     `json:"payment" binding:"required"`
}

type ItemModel struct {
	ID             *int
	OrderID        *int
	Product        *product.Model
	ProductVariant *product.VariantModel
	Quantity       *int
	Price          *int64
}
