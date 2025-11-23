package service

import (
	"time"

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
		Category:    &category,
	}
	if err := p.validate.Struct(product); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return product, nil
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
	price int64,
	quantity int,
) (*domain.ProductVariant, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	productVariant := &domain.ProductVariant{
		ID:            id,
		SKU:           sku,
		Price:         price,
		Quantity:      quantity,
		PurchaseCount: 0,
	}
	if err := p.validate.Struct(productVariant); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return productVariant, nil
}

func (p *Product) Update(
	product *domain.Product,
	name *string,
	description *string,
	Category *domain.Category,
) error {
	if name != nil {
		product.Name = *name
	}
	if description != nil {
		product.Description = *description
	}
	if Category != nil {
		product.Category = Category
	}
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) UpdateVariant(
	productVariant *domain.ProductVariant,
	price *int64,
	quantity *int,
) error {
	if price != nil {
		productVariant.Price = *price
	}
	if quantity != nil {
		productVariant.Quantity = *quantity
	}
	if err := p.validate.Struct(productVariant); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) UpdateOption(
	option *domain.Option,
	name *string,
) error {
	if name != nil {
		option.Name = *name
	}
	if err := p.validate.Struct(option); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) UpdateOptionValue(
	optionValue *domain.OptionValue,
	value *string,
) error {
	if value != nil {
		optionValue.Value = *value
	}
	if err := p.validate.Struct(optionValue); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) AddImage(product *domain.Product, image domain.ProductImage) error {
	product.Images = append(product.Images, image)
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) AddVariant(product *domain.Product, variant domain.ProductVariant) error {
	product.Variants = append(product.Variants, variant)
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) Remove(product *domain.Product) error {
	now := time.Now()
	if product.DeletedAt == nil {
		product.DeletedAt = &now
	}
	for _, option := range product.Options {
		if err := p.RemoveOptionsAndOptionValue(&option); err != nil {
			return err
		}
	}
	for i := range product.Variants {
		if err := p.RemoveVariant(&product.Variants[i]); err != nil {
			return err
		}
	}
	for i := range product.Images {
		if err := p.RemoveImage(&product.Images[i]); err != nil {
			return err
		}
	}
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) RemoveVariant(variant *domain.ProductVariant) error {
	now := time.Now()
	if variant.DeletedAt == nil {
		variant.DeletedAt = &now
	}
	if err := p.validate.Struct(variant); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) RemoveImage(image *domain.ProductImage) error {
	now := time.Now()
	if image.DeletedAt == nil {
		image.DeletedAt = &now
	}
	if err := p.validate.Struct(image); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) RemoveOptionsAndOptionValue(option *domain.Option) error {
	now := time.Now()
	if option.DeletedAt == nil {
		option.DeletedAt = &now
	}
	for _, v := range option.Values {
		if v.DeletedAt == nil {
			v.DeletedAt = &now
		}
	}
	if err := p.validate.Struct(option); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
