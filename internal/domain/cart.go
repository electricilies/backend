package domain

type Cart struct {
	ID    int
	Items *[]CartItem
}

type CartItem struct {
	ID             string
	ProductVariant *ProductVariant
	Quantity       int
}
