package domain

import "github.com/google/uuid"

type Cart struct {
	ID    int         `json:"id" binding:"required"`
	Items *[]CartItem `json:"items" binding:"omitnil,dive"`
}

type CartItem struct {
	ID             uuid.UUID       `json:"id" binding:"required,uuid"`
	ProductVariant *ProductVariant `json:"productVariant" binding:"omitnil"`
	Quantity       int             `json:"quantity" binding:"required"`
}
