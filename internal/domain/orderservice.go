package domain

import "github.com/google/uuid"

type OrderService interface {
	Create(
		userID uuid.UUID,
		address string,
		provider OrderProvider,
		totalAmount int64,
		items []OrderItem,
	) (*Order, error)

	CreateItem(
		productVariantID uuid.UUID,
		quantity int,
		price int64,
	) (*OrderItem, error)

	Update(
		order *Order,
		status *OrderStatus,
		address *string,
	) error
}
