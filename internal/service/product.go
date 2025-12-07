package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Product struct {
	validate *validator.Validate
}

func ProvideProduct(
	validate *validator.Validate,
) *Product {
	return &Product{
		validate: validate,
	}
}

var _ domain.ProductService = (*Product)(nil)

func (p *Product) Validate(
	product domain.Product,
) error {
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

// CreateOptionsWithOptionValues asfasdfsd
func (p *Product) CreateOptionsWithOptionValues(
	optionsWithOptionValues map[string][]string,
) (*[]domain.Option, error) {
	options := make([]domain.Option, 0, len(optionsWithOptionValues))
	for name, values := range optionsWithOptionValues {
		option, err := domain.NewProductOption(name)
		if err != nil {
			return nil, err
		}
		optionValues, err := domain.CreateOptionValues(values)
		if err != nil {
			return nil, err
		}
		option.Values = optionValues
		options = append(options, *option)
	}
	return &options, nil
}

func (p *Product) FilterProductVariantsInProducts(
	products []domain.Product,
	productVariantIDs []uuid.UUID,
) (*[]domain.ProductVariant, error) {
	var filteredVariants []domain.ProductVariant
	variantIDSet := make(map[uuid.UUID]struct{})
	for _, id := range productVariantIDs {
		variantIDSet[id] = struct{}{}
	}
	for _, product := range products {
		for _, variant := range product.Variants {
			if _, exists := variantIDSet[variant.ID]; exists {
				filteredVariants = append(filteredVariants, variant)
			}
		}
	}
	return &filteredVariants, nil
}
