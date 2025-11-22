package domain

import "github.com/google/uuid"

type CartService interface {
	Create(
		userID uuid.UUID,
	) (*Cart, error)

	CreateItem(
		productVariantID uuid.UUID,
		quantity int,
	) (*CartItem, error)

	AddItem(
		cart *Cart,
		item CartItem,
	) error

	UpdateItem(
		cart *Cart,
		itemID uuid.UUID,
		quantity int,
	) error

	RemoveItem(
		cart *Cart,
		itemID uuid.UUID,
	) error
}
