package domain

import "github.com/google/uuid"

type Cart struct {
	ID    int         `json:"id"    binding:"required"           validate:"required"`
	Items *[]CartItem `json:"items" validate:"omitnil,gt=0,dive"`
}

type CartItem struct {
	ID             uuid.UUID       `json:"id"             binding:"required" validate:"required,uuid"`
	ProductVariant *ProductVariant `json:"productVariant" validate:"omitnil"`
	Quantity       int             `json:"quantity"       binding:"required" validate:"required,gt=0,lte=100"`
}
