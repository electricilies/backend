package domain

import (
	"time"
)

type Product struct {
	ID            int               `json:"id" binding:"required"`
	Name          string            `json:"name" binding:"required"`
	Description   string            `json:"description" binding:"required"`
	ViewsCount    int               `json:"viewsCount" binding:"required"`
	TotalPurchase int               `json:"totalPurchase" binding:"required"`
	TrendingScore int64             `json:"trendingScore" binding:"required"`
	Price         int64             `json:"price" binding:"required"`
	Rating        float64           `json:"rating" binding:"required,gte=0,lte=5"`
	Options       *[]ProductOption  `json:"options,omitnil"`
	Images        *[]ProductImage   `json:"images,omitnil"`
	CreatedAt     time.Time         `json:"createdAt" binding:"required"`
	UpdatedAt     time.Time         `json:"updatedAt" binding:"required"`
	DeletedAt     time.Time         `json:"deletedAt,omitempty"`
	Category      *Category         `json:"category,omitnil"`
	Attributes    *[]Attribute      `json:"attributes,omitnil"`
	Variants      *[]ProductVariant `json:"variants,omitnil"`
}

type ProductOption struct {
	ID     int            `json:"id" binding:"required"`
	Name   string         `json:"name" binding:"required"`
	Values *[]OptionValue `json:"values,omitnil"`
}

type OptionValue struct {
	ID    int    `json:"id" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ProductVariant struct {
	ID            int             `json:"id" binding:"required"`
	SKU           string          `json:"sku" binding:"required"`
	Price         int64           `json:"price" binding:"required"`
	Quantity      int             `json:"quantity" binding:"required"`
	PurchaseCount int             `json:"purchaseCount" binding:"required"`
	CreatedAt     time.Time       `json:"createdAt" binding:"required"`
	DeletedAt     time.Time       `json:"deletedAt,omitempty"`
	OptionValues  *[]OptionValue  `json:"optionValues,omitnil"`
	Images        *[]ProductImage `json:"images,omitnil"`
	Product       *Product        `json:"product,omitempty"`
}

type ProductImage struct {
	ID        int       `json:"id" binding:"required"`
	URL       string    `json:"url" binding:"required,url"`
	Order     int       `json:"order" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
}
