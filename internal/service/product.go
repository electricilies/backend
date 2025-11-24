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

func (p *Product) Validate(
	product domain.Product,
) error {
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) Create(
	name string,
	description string,
	category domain.Category,
) (*domain.Product, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	now := time.Now()
	product := &domain.Product{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
		Category:    &category,
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

func (p *Product) CreateOptions(
	names []string,
) (*[]domain.Option, error) {
	options := make([]domain.Option, 0, len(names))
	for _, name := range names {
		option, err := p.CreateOption(name)
		if err != nil {
			return nil, err
		}
		options = append(options, *option)
	}
	return &options, nil
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

func (p *Product) CreateOptionValues(
	values []string,
) (*[]domain.OptionValue, error) {
	optionValues := make([]domain.OptionValue, 0, len(values))
	for _, v := range values {
		optionValue, err := p.CreateOptionValue(v)
		if err != nil {
			return nil, err
		}
		optionValues = append(optionValues, *optionValue)
	}
	return &optionValues, nil
}

func (p *Product) CreateOptionsWithOptionValues(
	optionsWithOptionValues map[string][]string,
) (*[]domain.Option, error) {
	options := make([]domain.Option, 0, len(optionsWithOptionValues))
	for name, values := range optionsWithOptionValues {
		option, err := p.CreateOption(name)
		if err != nil {
			return nil, err
		}
		optionValues, err := p.CreateOptionValues(values)
		if err != nil {
			return nil, err
		}
		option.Values = *optionValues
		options = append(options, *option)
	}
	return &options, nil
}

func (p *Product) AddOptions(product *domain.Product, option ...domain.Option) error {
	product.Options = append(product.Options, option...)
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) AddOptionValues(option *domain.Option, optionValue ...domain.OptionValue) error {
	option.Values = append(option.Values, optionValue...)
	if err := p.validate.Struct(option); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) AddVariants(product *domain.Product, variant ...domain.ProductVariant) error {
	product.Variants = append(product.Variants, variant...)
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) AddImages(product *domain.Product, image ...domain.ProductImage) error {
	product.Images = append(product.Images, image...)
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) AddVariantImages(product *domain.Product, variantID uuid.UUID, image ...domain.ProductImage) error {
	var variant *domain.ProductVariant
	for i := range product.Variants {
		if product.Variants[i].ID == variantID {
			variant = &product.Variants[i]
			break
		}
	}
	if variant == nil {
		return multierror.Append(domain.ErrNotFound, nil)
	}
	variant.Images = append(variant.Images, image...)
	if err := p.validate.Struct(variant); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
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
		ID:        id,
		URL:       url,
		Order:     order,
		CreatedAt: time.Now(),
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
		CreatedAt:     time.Now(),
	}
	if err := p.validate.Struct(productVariant); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return productVariant, nil
}

func (p *Product) AddAttributeValues(
	product *domain.Product,
	attributeValues ...domain.AttributeValue,
) error {
	product.AttributeValues = append(product.AttributeValues, attributeValues...)
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) Update(
	product *domain.Product,
	name *string,
	description *string,
	Category *domain.Category,
) error {
	updated := false
	if name != nil {
		product.Name = *name
	}
	if description != nil {
		product.Description = *description
	}
	if Category != nil {
		product.Category = Category
		updated = true
	}
	if updated {
		product.UpdatedAt = time.Now()
	}
	if err := p.validate.Struct(product); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) UpdateVariant(
	product *domain.Product,
	variantID uuid.UUID,
	price *int64,
	quantity *int,
) error {
	productVariant := product.GetVariantByID(variantID)
	update := false
	if price != nil {
		productVariant.Price = *price
		update = true
	}
	if quantity != nil {
		productVariant.Quantity = *quantity
		update = true
	}
	if !update {
		return nil
	}
	productVariant.UpdatedAt = time.Now()
	if err := p.validate.Struct(productVariant); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) UpdateOption(
	product *domain.Product,
	optionID uuid.UUID,
	name *string,
) error {
	option := product.GetOptionByID(optionID)
	if option == nil {
		return multierror.Append(domain.ErrNotFound, nil)
	}
	if name != nil {
		option.Name = *name
	}
	if err := p.validate.Struct(option); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) UpdateOptionValue(
	product *domain.Product,
	optionID uuid.UUID,
	optionValueID uuid.UUID,
	value *string,
) error {
	option := product.GetOptionByID(optionID)
	if option == nil {
		return multierror.Append(domain.ErrNotFound, nil)
	}
	optionValue := option.GetValueByID(optionValueID)
	if optionValue == nil {
		return multierror.Append(domain.ErrNotFound, nil)
	}
	if value != nil {
		optionValue.Value = *value
	}
	if err := p.validate.Struct(optionValue); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) Remove(product *domain.Product) error {
	now := time.Now()
	if product.DeletedAt == nil {
		product.UpdatedAt = now
		product.DeletedAt = &now
	}
	for _, option := range product.Options {
		if err := p.RemoveOptionsAndOptionValues(&option); err != nil {
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
		variant.UpdatedAt = now
	}
	if err := p.validate.Struct(variant); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (p *Product) RemoveVariants(variants ...*domain.ProductVariant) error {
	for _, variant := range variants {
		err := p.RemoveVariant(variant)
		if err != nil {
			return err
		}
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

func (p *Product) RemoveImages(images ...*domain.ProductImage) error {
	for _, image := range images {
		err := p.RemoveImage(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Product) RemoveOptionsAndOptionValues(options ...*domain.Option) error {
	now := time.Now()
	for _, option := range options {
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
	}
	return nil
}
