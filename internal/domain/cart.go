package domain

import "github.com/google/uuid"

type Cart struct {
	ID     uuid.UUID  `json:"id"     binding:"required"        validate:"required"`
	Items  []CartItem `json:"items"  validate:"omitempty,dive"`
	UserID uuid.UUID  `json:"userId" binding:"required"        validate:"required,uuid"`
}

type CartItem struct {
	ID               uuid.UUID `json:"id"                         binding:"required"        validate:"required,uuid"`
	ProductID        uuid.UUID `json:"productId"                  binding:"required"        validate:"required,uuid"`
	ProductVariantID uuid.UUID `json:"productVariantId,omitempty" validate:"omitempty,uuid"`
	Quantity         int       `json:"quantity"                   binding:"required"        validate:"required,gt=0,lte=100"`
}

func (c *Cart) AddItems(items ...CartItem) {
	if c.Items == nil {
		c.Items = []CartItem{}
	}
	c.Items = append(c.Items, items...)
}
