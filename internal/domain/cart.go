package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `json:"id"        binding:"required"                          validate:"required"`
	Items     []CartItem `json:"items"     validate:"omitempty,unique_cart_items,dive"`
	UserID    uuid.UUID  `json:"userId"    binding:"required"                          validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"                          validate:"required"`
}

type CartItem struct {
	ID               uuid.UUID `json:"id"                         binding:"required"   validate:"required"`
	ProductID        uuid.UUID `json:"productId"                  binding:"required"   validate:"required"`
	ProductVariantID uuid.UUID `json:"productVariantId,omitempty" validate:"omitempty"`
	Quantity         int       `json:"quantity"                   binding:"required"   validate:"required,gt=0,lte=100"`
}

func NewCart(userID uuid.UUID) (Cart, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return Cart{}, err
	}
	cart := Cart{
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
) (CartItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return CartItem{}, err
	}
	cartItem := CartItem{
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
