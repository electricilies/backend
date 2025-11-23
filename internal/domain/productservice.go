package domain

import "github.com/google/uuid"

type ProductService interface {
	Create(
		name string,
		description string,
		category Category,
	) (*Product, error)

	CreateOption(
		name string,
	) (*Option, error)

	CreateOptions(
		names []string,
	) (*[]Option, error)

	CreateOptionValues(
		values []string,
	) (*[]OptionValue, error)

	CreateOptionsWithOptionValues(
		optionsWithOptionValues map[string][]string,
	) (*[]Option, error)

	CreateImage(
		url string,
		order int,
	) (*ProductImage, error)

	CreateVariant(
		sku string,
		price int64,
		quantity int,
	) (*ProductVariant, error)

	AddAttributeValues(
		product *Product,
		attributeValues ...AttributeValue,
	) error

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

	AddVariantImages(
		product *Product,
		variant uuid.UUID,
		images ...ProductImage,
	) error

	Update(
		product *Product,
		name *string,
		description *string,
		category *Category,
	) error

	UpdateOption(
		product *Product,
		optionID uuid.UUID,
		name *string,
	) error

	UpdateOptionValue(
		product *Product,
		optionID uuid.UUID,
		optionValueID uuid.UUID,
		value *string,
	) error

	UpdateVariant(
		Product *Product,
		VariantID uuid.UUID,
		price *int64,
		quantity *int,
	) error

	Remove(
		product *Product,
	) error

	RemoveVariant(
		variant *ProductVariant,
	) error

	RemoveVariants(
		variants ...*ProductVariant,
	) error

	RemoveImage(
		image *ProductImage,
	) error

	RemoveImages(
		images ...*ProductImage,
	) error

	RemoveOptionsAndOptionValues(
		options ...*Option,
	) error
}
