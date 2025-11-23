package domain

type ProductService interface {
	Create(
		name string,
		description string,
		category Category,
	) (*Product, error)

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

	AddImage(
		product *Product,
		images ProductImage,
	) error

	AddVariant(
		product *Product,
		variant ProductVariant,
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
