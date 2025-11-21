package domain

type OrderService interface {
	Create(
		Address string,
		Provider OrderProvider,
		Status OrderStatus,
		IsPaid bool,
		TotalAmount int64,
	) (*Order, error)

	CreateItem(
		productVariant ProductVariant,
		quantity int,
		price int64,
	) (*OrderItem, error)
}
