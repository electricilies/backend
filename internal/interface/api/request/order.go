package request

type OrderStatus string

type CreateOrder struct {
	UserID string      `json:"userId" binding:"required"`
	Items  []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
	ProductID        int   `json:"productId" binding:"required"`
	ProductVariantID int   `json:"productVariantId,omitempty"`
	Quantity         int   `json:"quantity" binding:"required"`
	Price            int64 `json:"price" binding:"required"`
}

type UpdateOrderStatus struct {
	OrderStatus OrderStatus `json:"orderStatus" binding:"required"`
}
