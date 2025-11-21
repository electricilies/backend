package domain

type ProductService interface {
	Create(
		name string,
		description string,
		category Category,
	) (*Product, error)

	CreateOption(
		name string,
	) (*Option, error)

	CreateImage(
		url string,
		order int,
	) (*ProductImage, error)

	CreateVariant(
		sku string,
		price int64,
		quantity int,
	) (*ProductVariant, error)

	CreateOptionValue(
		value string,
	) (*OptionValue, error)
}
