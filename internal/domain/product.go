package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Product struct {
	ID                uuid.UUID        `validate:"required"`
	Name              string           `validate:"required,gte=3,lte=200"`
	Description       string           `validate:"required,gte=10"`
	ViewsCount        int              `validate:"gte=0"`
	TotalPurchase     int              `validate:"gte=0"`
	TrendingScore     int64            `validate:"gte=0"`
	Price             int64            `validate:"required,gt=0"`
	Rating            float64          `validate:"gte=0,lte=5"`
	Options           []Option         `validate:"omitempty,unique=ID,unique=Name,dive"`
	Images            []ProductImage   `validate:"gt=0,dive"`
	CreatedAt         time.Time        `validate:"required"`
	UpdatedAt         time.Time        `validate:"required,gtefield=CreatedAt"`
	DeletedAt         time.Time        `validate:"omitempty,gtefield=CreatedAt"`
	CategoryID        uuid.UUID        `validate:"required"`
	AttributeIDs      []uuid.UUID      `validate:"omitempty,dive,required"`
	AttributeValueIDs []uuid.UUID      `validate:"omitempty,dive,required"`
	Variants          []ProductVariant `validate:"gt=0,unique=ID,unique=SKU,productVariantStructure,dive"`
}

type Option struct {
	ID        uuid.UUID     `validate:"required"`
	Name      string        `validate:"required"`
	Values    []OptionValue `validate:"omitempty,unique=ID,unique=Value,dive"`
	DeletedAt time.Time     `validate:"omitempty"`
}

type OptionValue struct {
	ID        uuid.UUID `validate:"required"`
	Value     string    `validate:"required"`
	DeletedAt time.Time `validate:"omitempty"`
}

type ProductVariant struct {
	ID            uuid.UUID      `validate:"required"`
	SKU           string         `validate:"required"`
	Price         int64          `validate:"required,gt=0"`
	Quantity      int            `validate:"gte=0"`
	PurchaseCount int            `validate:"gte=0"`
	CreatedAt     time.Time      `validate:"required"`
	UpdatedAt     time.Time      `validate:"required,gtefield=CreatedAt"`
	DeletedAt     time.Time      `validate:"omitempty,gtefield=CreatedAt"`
	OptionValues  []OptionValue  `validate:"omitempty,unique=ID,unique=Value,dive"`
	Images        []ProductImage `validate:"omitempty,unique=ID,unique=URL,unique=Order,dive"`
}

type ProductImage struct {
	ID        uuid.UUID `validate:"required"`
	URL       string    `validate:"required,url"`
	Order     int       `validate:"required,gte=0"`
	CreatedAt time.Time `validate:"required"`
	DeletedAt time.Time `validate:"omitempty,gtefield=CreatedAt"`
}

func NewProduct(
	name string,
	description string,
	categoryID uuid.UUID,
) (*Product, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(ErrInternal, err)
	}
	now := time.Now()
	product := &Product{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
		CategoryID:  categoryID,
	}
	return product, nil
}

func NewProductOption(
	name string,
) (*Option, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(ErrInternal, err)
	}
	option := &Option{
		ID:   id,
		Name: name,
	}
	return option, nil
}

func NewProductImage(
	url string,
	order int,
) (*ProductImage, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(ErrInternal, err)
	}
	productImage := &ProductImage{
		ID:        id,
		URL:       url,
		Order:     order,
		CreatedAt: time.Now(),
	}
	return productImage, nil
}

func NewVariant(
	sku string,
	price int64,
	quantity int,
) (*ProductVariant, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(ErrInternal, err)
	}
	now := time.Now()
	productVariant := &ProductVariant{
		ID:            id,
		SKU:           sku,
		Price:         price,
		Quantity:      quantity,
		PurchaseCount: 0,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	return productVariant, nil
}

func (p *Product) AddAttributeIDs(attributeIDs ...uuid.UUID) {
	p.AttributeIDs = append(p.AttributeIDs, attributeIDs...)
}

func (p *Product) AddAttributeValueIDs(attributeValueIDs ...uuid.UUID) {
	p.AttributeValueIDs = append(p.AttributeValueIDs, attributeValueIDs...)
}

func (p *Product) AddOptions(options ...Option) {
	p.Options = append(p.Options, options...)
}

func (p *Product) AddVariants(variants ...ProductVariant) {
	p.Variants = append(p.Variants, variants...)
}

func (p *Product) AddImages(images ...ProductImage) {
	p.Images = append(p.Images, images...)
}

func (p *Product) AddVariantImages(variantID uuid.UUID, images ...ProductImage) error {
	var variant *ProductVariant
	for i := range p.Variants {
		if p.Variants[i].ID == variantID {
			variant = &p.Variants[i]
			break
		}
	}
	if variant == nil {
		return multierror.Append(ErrNotFound, nil)
	}
	variant.Images = append(variant.Images, images...)
	return nil
}

func (p *Product) Update(
	name string,
	description string,
	categoryID uuid.UUID,
) {
	updated := false
	if name != "" && p.Name != name {
		p.Name = name
		updated = true
	}
	if description != "" && p.Description != description {
		p.Description = description
		updated = true
	}
	if categoryID != uuid.Nil {
		p.CategoryID = categoryID
		updated = true
	}
	if updated {
		p.UpdatedAt = time.Now()
	}
}

func (p *Product) UpdateVariant(
	variantID uuid.UUID,
	price int64,
	quantity int,
) error {
	var variant *ProductVariant
	for i := range p.Variants {
		if p.Variants[i].ID == variantID {
			variant = &p.Variants[i]
			break
		}
	}
	if variant == nil {
		return multierror.Append(ErrNotFound, nil)
	}
	updated := false
	if price != 0 && variant.Price != price {
		variant.Price = price
		updated = true
	}
	if quantity != 0 {
		variant.Quantity = quantity
		updated = true
	}
	if updated {
		variant.UpdatedAt = time.Now()
	}
	return nil
}

func (p *Product) UpdateOption(
	optionID uuid.UUID,
	name string,
) error {
	var option *Option
	for i := range p.Options {
		if p.Options[i].ID == optionID {
			option = &p.Options[i]
			break
		}
	}
	if option == nil {
		return multierror.Append(ErrNotFound, nil)
	}
	if name != "" && option.Name != name {
		option.Name = name
	}
	return nil
}

func (p *Product) UpdateOptionValue(
	optionID uuid.UUID,
	optionValueID uuid.UUID,
	value string,
) error {
	var option *Option
	for i := range p.Options {
		if p.Options[i].ID == optionID {
			option = &p.Options[i]
			break
		}
	}
	if option == nil {
		return multierror.Append(ErrNotFound, nil)
	}
	var optionValue *OptionValue
	for i := range option.Values {
		if option.Values[i].ID == optionValueID {
			optionValue = &option.Values[i]
			break
		}
	}
	if optionValue == nil {
		return multierror.Append(ErrNotFound, nil)
	}
	if value != "" && optionValue.Value != value {
		optionValue.Value = value
	}
	return nil
}

func (p *Product) GetOptionByID(optionID uuid.UUID) *Option {
	for _, option := range p.Options {
		if option.ID == optionID {
			return &option
		}
	}
	return nil
}

func (p *Product) GetOptionsByIDs(optionIDs []uuid.UUID) []Option {
	var options []Option
	optionIDSet := make(map[uuid.UUID]struct{})
	for _, id := range optionIDs {
		optionIDSet[id] = struct{}{}
	}
	for _, option := range p.Options {
		if _, exists := optionIDSet[option.ID]; exists {
			options = append(options, option)
		}
	}
	return options
}

func (p *Product) GetVariantByID(variantID uuid.UUID) *ProductVariant {
	for _, variant := range p.Variants {
		if variant.ID == variantID {
			return &variant
		}
	}
	return nil
}

func (p *Product) UpdateMinPrice() {
	if len(p.Variants) == 0 {
		return
	}
	minPrice := p.Variants[0].Price
	for _, variant := range p.Variants {
		if variant.Price < minPrice {
			minPrice = variant.Price
		}
	}
	p.Price = minPrice
}

func (p *Product) Remove() {
	now := time.Now()
	p.UpdatedAt = now
	p.DeletedAt = now
	for i := range p.Options {
		p.Options[i].Remove()
	}
	for i := range p.Variants {
		p.Variants[i].Remove()
	}
	for i := range p.Images {
		p.Images[i].Remove()
	}
}

func (o *Option) AddOptionValues(optionValues ...OptionValue) {
	o.Values = append(o.Values, optionValues...)
}

func (o *Option) Remove() {
	now := time.Now()
	o.DeletedAt = now
	for i := range o.Values {
		o.Values[i].Remove()
	}
}

func (ov *OptionValue) Remove() {
	now := time.Now()
	ov.DeletedAt = now
}

func (v *ProductVariant) Remove() {
	now := time.Now()
	v.DeletedAt = now
	v.UpdatedAt = now
}

func (img *ProductImage) Remove() {
	now := time.Now()
	img.DeletedAt = now
}

func CreateOptionValues(
	values []string,
) ([]OptionValue, error) {
	optionValues := make([]OptionValue, 0, len(values))
	for _, value := range values {
		id, err := uuid.NewV7()
		if err != nil {
			return nil, multierror.Append(ErrInternal, err)
		}
		optionValue := OptionValue{
			ID:    id,
			Value: value,
		}
		optionValues = append(optionValues, optionValue)
	}
	return optionValues, nil
}

func (o *Option) GetValueByID(optionValueID uuid.UUID) *OptionValue {
	for _, value := range o.Values {
		if value.ID == optionValueID {
			return &value
		}
	}
	return nil
}

func (o *Option) GetValuesByIDs(optionValueIDs []uuid.UUID) []OptionValue {
	var values []OptionValue
	optionValueIDSet := make(map[uuid.UUID]struct{})
	for _, id := range optionValueIDs {
		optionValueIDSet[id] = struct{}{}
	}
	for _, value := range o.Values {
		if _, exists := optionValueIDSet[value.ID]; exists {
			values = append(values, value)
		}
	}
	return values
}
