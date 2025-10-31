package response

import "time"

type Order struct {
	ID            int         `json:"id"`
	UserID        string      `json:"user_id"`
	OrderStatusID int         `json:"order_status_id"`
	PaymentID     int         `json:"payment_id"`
	CreatedAt     *time.Time  `json:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
	Items         []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID               int     `json:"id"`
	OrderID          int     `json:"order_id"`
	ProductID        int     `json:"product_id"`
	ProductVariantID int     `json:"product_variant_id"`
	ProductName      string  `json:"product_name"`
	Quantity         int     `json:"quantity"`
	Price            float64 `json:"price"`
}
