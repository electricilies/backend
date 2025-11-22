package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID         `json:"id"            binding:"required"                      validate:"required"`
	Name          string            `json:"name"          binding:"required"                      validate:"required,gte=3,lte=200"`
	Description   string            `json:"description"   binding:"required"                      validate:"required,gte=10"`
	ViewsCount    int               `json:"viewsCount"    binding:"required"                      validate:"required,gte=0"`
	TotalPurchase int               `json:"totalPurchase" binding:"required"                      validate:"required,gte=0"`
	TrendingScore int64             `json:"trendingScore" binding:"required"                      validate:"required,gte=0"`
	Price         int64             `json:"price"         binding:"required"                      validate:"required,gt=0"`
	Rating        float64           `json:"rating"        binding:"required"                      validate:"required,gte=0,lte=5"`
	Options       *[]Option         `json:"options"       validate:"omitnil,dive"`
	Images        *[]ProductImage   `json:"images"        validate:"omitnil,gte=1,dive"`
	CreatedAt     time.Time         `json:"createdAt"     binding:"required"                      validate:"required"`
	UpdatedAt     time.Time         `json:"updatedAt"     binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt     *time.Time        `json:"deletedAt"     validate:"omitempty,gtefield=CreatedAt"`
	Category      *Category         `json:"category"`
	Attributes    *[]Attribute      `json:"attributes"    validate:"omitnil,dive"`
	Variants      *[]ProductVariant `json:"variants"      validate:"omitnil,gte=1,dive"`
}

type Option struct {
	ID     uuid.UUID      `json:"id"     binding:"required"           validate:"required"`
	Name   string         `json:"name"   binding:"required"           validate:"required"`
	Values *[]OptionValue `json:"values" validate:"omitnil,gt=0,dive"`
}

type OptionValue struct {
	ID        uuid.UUID  `json:"id"        binding:"required"   validate:"required"`
	Value     string     `json:"value"     binding:"required"   validate:"required"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty"`
}

type ProductVariant struct {
	ID            uuid.UUID       `json:"id"            binding:"required"                      validate:"required"`
	SKU           string          `json:"sku"           binding:"required"                      validate:"required"`
	Price         int64           `json:"price"         binding:"required"                      validate:"required,gt=0"`
	Quantity      int             `json:"quantity"      binding:"required"                      validate:"required,gte=0"`
	PurchaseCount int             `json:"purchaseCount" binding:"required"                      validate:"required,gte=0"`
	CreatedAt     time.Time       `json:"createdAt"     binding:"required"                      validate:"required"`
	DeletedAt     *time.Time      `json:"deletedAt"     validate:"omitempty,gtefield=CreatedAt"`
	OptionValues  *[]OptionValue  `json:"optionValues"  validate:"omitnil,gte=1,dive"`
	Images        *[]ProductImage `json:"images"        validate:"omitnil,dive"`
	Product       *Product        `json:"product"       validate:"omitnil"`
}

type ProductImage struct {
	ID        uuid.UUID  `json:"id"         binding:"required"                      validate:"required"`
	URL       string     `json:"url"        binding:"required"                      validate:"required,url"`
	Order     int        `json:"order"      binding:"required"                      validate:"required,gte=0"`
	CreatedAt time.Time  `json:"createdAt"  binding:"required"                      validate:"required"`
	DeletedAt *time.Time `json:"deletedAt " validate:"omitempty,gtefield=CreatedAt"`
}
