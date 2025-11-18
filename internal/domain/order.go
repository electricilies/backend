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
	ID        int         `json:"id" binding:"required"`
	User      User        `json:"user" binding:"required"`
	Address   string      `json:"address" binding:"required"`
	Status    OrderStatus `json:"status" binding:"required"`
	CreatedAt time.Time   `json:"createdAt" binding:"required"`
	UpdatedAt time.Time   `json:"updatedAt" binding:"required"`
	Items     []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
	ID             int            `json:"id" binding:"required"`
	OrderID        int            `json:"orderId" binding:"required"`
	ProductVariant ProductVariant `json:"productVariant" binding:"required"`
	Quantity       int            `json:"quantity" binding:"required"`
	Price          int64          `json:"price" binding:"required"`
}
