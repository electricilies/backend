package domain

type CreateOrderParam struct {
	UserID string          `binding:"required"`
	Data   CreateOrderData `binding:"required"`
}

type CreateOrderData struct {
	Address string                `json:"address" binding:"required"`
	Items   []CreateOrderItemData `json:"items" binding:"required,dive"`
}

type CreateOrderItemData struct {
	ProductVariantID int   `json:"productVariantId" binding:"required"`
	Quantity         int   `json:"quantity" binding:"required"`
	Price            int64 `json:"price" binding:"required"`
}

type UpdateOrderParam struct {
	OrderID int             `binding:"required"`
	Data    UpdateOrderData `binding:"required"`
}

type UpdateOrderData struct {
	Status  *OrderStatus `json:"status" binding:"omitnil,oneof=Pending Processing Shipped Delivered Cancelled"`
	Address *string      `json:"address" binding:"omitnil"`
}

type GetOrderParam struct {
	OrderID int `binding:"required"`
}

type DeleteOrderParam struct {
	OrderID int `binding:"required"`
}
