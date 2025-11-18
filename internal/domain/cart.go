package domain

import "github.com/google/uuid"

type Cart struct {
	ID    int         `json:"id" binding:"required"`
	Items *[]CartItem `json:"items,omitnil"`
}

type CartItem struct {
	ID             uuid.UUID       `json:"id" binding:"required"`
	ProductVariant *ProductVariant `json:"productVariant,omitnil"`
	Quantity       int             `json:"quantity" binding:"required"`
}
