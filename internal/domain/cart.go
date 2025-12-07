package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `validate:"required"`
	Items     []CartItem `validate:"omitempty,unique=ID,unique=ProductVariantID,dive"`
	UserID    uuid.UUID  `validate:"required"`
	UpdatedAt time.Time  `validate:"required"`
}

type CartItem struct {
	ID               uuid.UUID `validate:"required"`
	ProductID        uuid.UUID `validate:"required"`
	ProductVariantID uuid.UUID `validate:"omitempty"`
	Quantity         int       `validate:"required,gt=0,lte=100"`
}

func NewCart(userID uuid.UUID) (*Cart, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	cart := &Cart{
		ID:        id,
		UserID:    userID,
		Items:     []CartItem{},
		UpdatedAt: time.Now(),
	}
	return cart, nil
}

func NewCartItem(
	productID uuid.UUID,
	productVariantID uuid.UUID,
	quantity int,
) (*CartItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	cartItem := &CartItem{
		ID:               id,
		ProductID:        productID,
		ProductVariantID: productVariantID,
		Quantity:         quantity,
	}
	return cartItem, nil
}

func (c *Cart) UpsertItem(item CartItem) CartItem {
	for i := range c.Items {
		existingItem := &c.Items[i]
		if existingItem.ProductID == item.ProductID {
			existingItem.Quantity += item.Quantity
			return *existingItem
		}
	}
	c.Items = append(c.Items, item)
	return item
}

func (c *Cart) RemoveItem(itemID uuid.UUID) {
	filteredItems := []CartItem{}
	for _, item := range c.Items {
		if item.ID != itemID {
			filteredItems = append(filteredItems, item)
		}
	}
	c.Items = filteredItems
}

func (c *Cart) UpdateItem(
	itemID uuid.UUID,
	quantity int,
) {
	for i, item := range c.Items {
		if item.ID == itemID {
			if quantity == 0 {
				c.RemoveItem(itemID)
			} else {
				c.Items[i].Quantity = quantity
			}
			break
		}
	}
}

func (c *Cart) ClearItems() {
	c.Items = []CartItem{}
}
