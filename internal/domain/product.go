package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID              uuid.UUID        `json:"id"            binding:"required"                      validate:"required"`
	Name            string           `json:"name"          binding:"required"                      validate:"required,gte=3,lte=200"`
	Description     string           `json:"description"   binding:"required"                      validate:"required,gte=10"`
	ViewsCount      int              `json:"viewsCount"    binding:"required"                      validate:"required,gte=0"`
	TotalPurchase   int              `json:"totalPurchase" binding:"required"                      validate:"required,gte=0"`
	TrendingScore   int64            `json:"trendingScore" binding:"required"                      validate:"required,gte=0"`
	Price           int64            `json:"price"         binding:"required"                      validate:"required,gt=0"`
	Rating          float64          `json:"rating"        binding:"required"                      validate:"required,gte=0,lte=5"`
	Options         []Option         `json:"options"       validate:"omitempty,dive"`
	Images          []ProductImage   `json:"images"        validate:"omitempty,dive"`
	CreatedAt       time.Time        `json:"createdAt"     binding:"required"                      validate:"required"`
	UpdatedAt       time.Time        `json:"updatedAt"     binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt       *time.Time       `json:"deletedAt"     validate:"omitempty,gtefield=CreatedAt"`
	Category        *Category        `json:"category"`
	AttributeValues []AttributeValue `json:"attributes"    validate:"omitempty,dive"`
	Variants        []ProductVariant `json:"variants"      validate:"omitempty,dive"`
}

type Option struct {
	ID        uuid.UUID     `json:"id"        binding:"required"        validate:"required"`
	Name      string        `json:"name"      binding:"required"        validate:"required"`
	Values    []OptionValue `json:"values"    validate:"omitempty,dive"`
	DeletedAt *time.Time    `json:"deletedAt" validate:"omitempty"`
}

type OptionValue struct {
	ID        uuid.UUID  `json:"id"        binding:"required"   validate:"required"`
	Value     string     `json:"value"     binding:"required"   validate:"required"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty"`
}

type ProductVariant struct {
	ID            uuid.UUID      `json:"id"                binding:"required"                      validate:"required"`
	SKU           string         `json:"sku"               binding:"required"                      validate:"required"`
	Price         int64          `json:"price"             binding:"required"                      validate:"required,gt=0"`
	Quantity      int            `json:"quantity"          binding:"required"                      validate:"required,gte=0"`
	PurchaseCount int            `json:"purchaseCount"     binding:"required"                      validate:"required,gte=0"`
	CreatedAt     time.Time      `json:"createdAt"         binding:"required"                      validate:"required"`
	UpdatedAt     time.Time      `json:"updatedAt"         binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt     *time.Time     `json:"deletedAt"         validate:"omitempty,gtefield=CreatedAt"`
	OptionValues  []OptionValue  `json:"optionValues"      validate:"omitempty,dive"`
	Images        []ProductImage `json:"images"            validate:"omitempty,dive"`
	Product       *Product       `json:"product,omitempty"`
}

type ProductImage struct {
	ID        uuid.UUID  `json:"id"         binding:"required"                      validate:"required"`
	URL       string     `json:"url"        binding:"required"                      validate:"required,url"`
	Order     int        `json:"order"      binding:"required"                      validate:"required,gte=0"`
	CreatedAt time.Time  `json:"createdAt"  binding:"required"                      validate:"required"`
	DeletedAt *time.Time `json:"deletedAt " validate:"omitempty,gtefield=CreatedAt"`
}

func (p *Product) GetOptionByID(optionID uuid.UUID) *Option {
	for _, option := range p.Options {
		if option.ID == optionID {
			return &option
		}
	}
	return nil
}

func (p *Product) GetOptionsByIDs(optionIDs []uuid.UUID) []*Option {
	var options []*Option
	optionIDSet := make(map[uuid.UUID]struct{})
	for _, id := range optionIDs {
		optionIDSet[id] = struct{}{}
	}
	for _, option := range p.Options {
		if _, exists := optionIDSet[option.ID]; exists {
			options = append(options, &option)
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

func (o *Option) GetValueByID(optionValueID uuid.UUID) *OptionValue {
	for _, value := range o.Values {
		if value.ID == optionValueID {
			return &value
		}
	}
	return nil
}

func (o *Option) GetValuesByIDs(optionValueIDs []uuid.UUID) []*OptionValue {
	var values []*OptionValue
	optionValueIDSet := make(map[uuid.UUID]struct{})
	for _, id := range optionValueIDs {
		optionValueIDSet[id] = struct{}{}
	}
	for _, value := range o.Values {
		if _, exists := optionValueIDSet[value.ID]; exists {
			values = append(values, &value)
		}
	}
	return values
}
