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

	CreateOptionValue(
		value string,
	) (*OptionValue, error)

	CreateImage(
		url string,
		order int,
	) (*ProductImage, error)

	CreateVariant(
		sku string,
		price int64,
		quantity int,
	) (*ProductVariant, error)

	AddOptions(
		product *Product,
		options ...Option,
	) error

	AddOptionValues(
		option *Option,
		optionValues ...OptionValue,
	) error

	AddVariants(
		product *Product,
		variants ...ProductVariant,
	) error

	AddImages(
		product *Product,
		images ...ProductImage,
	) error

	Update(
		product *Product,
		name *string,
		description *string,
		category *Category,
	) error

	UpdateOption(
		option *Option,
		name *string,
	) error

	UpdateOptionValue(
		optionValue *OptionValue,
		value *string,
	) error

	UpdateVariant(
		variant *ProductVariant,
		price *int64,
		quantity *int,
	) error

	Remove(
		product *Product,
	) error

	RemoveVariant(
		variant *ProductVariant,
	) error

	RemoveImage(
		image *ProductImage,
	) error

	RemoveOptionsAndOptionValue(
		option *Option,
	) error
}
