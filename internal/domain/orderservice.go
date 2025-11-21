package domain

type OrderService interface {
	Create(
		address string,
		provider OrderProvider,
		isPaid bool,
		totalAmount int64,
	) (*Order, error)

	CreateItem(
		productVariant ProductVariant,
		quantity int,
		price int64,
	) (*OrderItem, error)
}
