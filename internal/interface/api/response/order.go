package response

import "time"

type Order struct {
	ID            int         `json:"id" binding:"required"`
	UserID        string      `json:"user_id" binding:"required"`
	OrderStatusID int         `json:"order_status_id" binding:"required"`
	PaymentID     int         `json:"payment_id" binding:"required"`
	CreatedAt     time.Time   `json:"created_at" binding:"required"`
	UpdatedAt     time.Time   `json:"updated_at" binding:"required"`
	Items         []OrderItem `json:"items," binding:"required"`
}

type OrderItem struct {
	ID             int            `json:"id" binding:"required"`
	OrderID        int            `json:"order_id" binding:"required"`
	Product        Product        `json:"product" binding:"required"`
	ProductVariant ProductVariant `json:"product_variant" binding:"required"`
	Quantity       int            `json:"quantity" binding:"required"`
	Price          float64        `json:"price" binding:"required"`
}
