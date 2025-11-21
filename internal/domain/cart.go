package domain

import "github.com/google/uuid"

type Cart struct {
	ID    uuid.UUID  `json:"id"    binding:"required"             validate:"required"`
	Items []CartItem `json:"items" validate:"omitempty,gt=0,dive"`
}

func (c *Cart) AddItems(items ...CartItem) {
	c.Items = append(c.Items, items...)
}

type CartItem struct {
	ID             uuid.UUID      `json:"id"             binding:"required" validate:"required,uuid"`
	ProductVariant ProductVariant `json:"productVariant" binding:"required" validate:"required"`
	Quantity       int            `json:"quantity"       binding:"required" validate:"required,gt=0,lte=100"`
}
