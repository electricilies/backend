package domain

import "time"

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "Pending"
	OrderStatusProcessing OrderStatus = "Processing"
	OrderStatusShipped    OrderStatus = "Shipped"
	OrderStatusDelivered  OrderStatus = "Delivered"
	OrderStatusCancelled  OrderStatus = "Cancelled"
)

type Order struct {
	ID          int          `json:"id" binding:"required" validate:"required"`
	User        *User        `json:"user"`
	Address     string       `json:"address" binding:"required" validate:"required"`
	Status      OrderStatus  `json:"status" binding:"required" validate:"required,oneof=Pending Processing Shipped Delivered Cancelled"`
	CreatedAt   time.Time    `json:"createdAt" binding:"required" validate:"required"`
	UpdatedAt   time.Time    `json:"updatedAt" binding:"required" validate:"required"`
	Items       *[]OrderItem `json:"items" binding:"omitnil" validate:"omitnil,min=1,dive"`
	TotalAmount int64        `json:"totalAmount" binding:"required" validate:"required"`
}

type OrderItem struct {
	ID             int             `json:"id" binding:"required" validate:"required"`
	OrderID        int             `json:"orderId" binding:"required" validate:"required"`
	ProductVariant *ProductVariant `json:"productVariant" binding:"omitnil"`
	Quantity       int             `json:"quantity" binding:"required" validate:"required,gt=0"`
	Price          int64           `json:"price" binding:"required" validate:"required,gt=0"`
}
