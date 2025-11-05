package request

type CreateOrder struct {
	UserID string             `json:"userId" binding:"required"`
	Items  []OrderItemRequest `json:"items" binding:"required"`
}

type OrderItemRequest struct {
	ProductID        int `json:"productId" binding:"required"`
	ProductVariantID int `json:"productVariantId,omitempty"`
	Quantity         int `json:"quantity" binding:"required"`
}

type UpdateOrderStatus struct {
	OrderStatusID int `json:"orderStatusId" binding:"required"`
}
