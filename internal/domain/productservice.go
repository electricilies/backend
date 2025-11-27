package domain

import "github.com/google/uuid"

type ProductService interface {
	Validate(
		product Product,
	) error

	CreateOptionsWithOptionValues(
		optionsWithOptionValues map[string][]string,
	) (*[]Option, error)

	FilterProductVariantsInProducts(
		products []Product,
		productVariantIDs []uuid.UUID,
	) (*[]ProductVariant, error)
}
