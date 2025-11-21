package domain

type CartService interface {
	CreateItem(
		productVariant ProductVariant,
		quantity int,
	) (*CartItem, error)
}
