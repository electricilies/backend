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

var _ domain.ProductService = &Product{}

func (p *Product) Create(
	name string,
	description string,
	category domain.Category,
) (*domain.Product, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	product := &domain.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Category:    category,
	}
	if err := p.validate.Struct(product); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return product, nil
}

func (p *Product) CreateOption(
	name string,
) (*domain.Option, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	option := &domain.Option{
		ID:   id,
		Name: name,
	}
	if err := p.validate.Struct(option); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return option, nil
}

func (p *Product) CreateImage(
	url string,
	order int,
) (*domain.ProductImage, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	productImage := &domain.ProductImage{
		ID:    id,
		URL:   url,
		Order: order,
	}
	if err := p.validate.Struct(productImage); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return productImage, nil
}

func (p *Product) CreateVariant(
	sku string,
	price int,
	quantity int,
) (*domain.ProductVariant, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	productVariant := &domain.ProductVariant{
		ID:            id,
		SKU:           sku,
		Price:         int64(price),
		Quantity:      quantity,
		PurchaseCount: 0,
	}
	if err := p.validate.Struct(productVariant); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return productVariant, nil
}

func (p *Product) CreateOptionValue(
	value string,
) (*domain.OptionValue, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	optionValue := &domain.OptionValue{
		ID:    id,
		Value: value,
	}
	if err := p.validate.Struct(optionValue); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return optionValue, nil
}
