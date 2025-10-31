package request

type CreateOrder struct {
	UserID    string             `json:"user_id" binding:"required"`
	PaymentID int                `json:"payment_id" binding:"required"`
	Items     []OrderItemRequest `json:"items" binding:"required"`
}

type OrderItemRequest struct {
	ProductID        int `json:"product_id" binding:"required"`
	ProductVariantID int `json:"product_variant_id,omitempty"`
	Quantity         int `json:"quantity" binding:"required"`
}

type UpdateOrderStatus struct {
	OrderStatusID int `json:"order_status_id" binding:"required"`
}
