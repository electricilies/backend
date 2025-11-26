package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `json:"id"        binding:"required"        validate:"required"`
	Items     []CartItem `json:"items"     validate:"omitempty,dive"`
	UserID    uuid.UUID  `json:"userId"    binding:"required"        validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"        validate:"required"`
}

type CartItem struct {
	ID               uuid.UUID `json:"id"                         binding:"required"   validate:"required"`
	ProductID        uuid.UUID `json:"productId"                  binding:"required"   validate:"required"`
	ProductVariantID uuid.UUID `json:"productVariantId,omitempty" validate:"omitempty"`
	Quantity         int       `json:"quantity"                   binding:"required"   validate:"required,gt=0,lte=100"`
}

func (c *Cart) AddItems(items ...CartItem) {
	if c.Items == nil {
		c.Items = []CartItem{}
	}
	c.Items = append(c.Items, items...)
}

func (c *Cart) RemoveItems(itemID ...uuid.UUID) {
	filteredItems := []CartItem{}
	itemIDMap := make(map[uuid.UUID]struct{})
	for _, id := range itemID {
		itemIDMap[id] = struct{}{}
	}

	for _, item := range c.Items {
		if _, found := itemIDMap[item.ID]; !found {
			filteredItems = append(filteredItems, item)
		}
	}

	c.Items = filteredItems
}

func (c *Cart) UpdateItem(
	itemID uuid.UUID,
	quantity *int,
) {
	for i, item := range c.Items {
		if item.ID == itemID {
			if quantity != nil {
				if *quantity == 0 {
					c.RemoveItems(itemID)
				} else {
					c.Items[i].Quantity = *quantity
				}
			}
			break
		}
	}
}
