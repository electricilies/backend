package service

import "backend/internal/domain"

type CreateOrderParam struct {
	UserID string
	Data   CreateOrderData
}

type CreateOrderData struct {
	Address string                `json:"address" binding:"required"`
	Items   []CreateOrderItemData `json:"items" binding:"required,dive,required"`
}

type CreateOrderItemData struct {
	ProductVariantID int   `json:"productVariantId" binding:"required"`
	Quantity         int   `json:"quantity" binding:"required"`
	Price            int64 `json:"price" binding:"required"`
}

type UpdateOrderParam struct {
	OrderID int
	Data    UpdateOrderData
}

type UpdateOrderData struct {
	Status  *domain.OrderStatus `json:"status" binding:"required,oneof=Pending Processing Shipped Delivered Cancelled"`
	Address *string             `json:"address" binding:"required"`
}

type ListOrderParam struct {
	PaginationParam
	OrderIDs  *[]int
	UserIDs   *[]string
	StatusIDs *[]int
}

type GetOrderParam struct {
	OrderID int
}

type DeleteOrderParam struct {
	OrderID int
}
