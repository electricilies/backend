package domain

import "github.com/google/uuid"

type CartService interface {
	Create(
		userID uuid.UUID,
	) (*Cart, error)

	CreateItem(
		productVariant ProductVariant,
		quantity int,
	) (*CartItem, error)
}
