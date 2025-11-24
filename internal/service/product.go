package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
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

var _ domain.ProductService = &Product{}

func (p *Product) Validate(
	product domain.Product,
) error {
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
