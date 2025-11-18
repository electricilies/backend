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
	Options       *[]ProductOption  `json:"options" binding:"omitempty,dive"`
	Images        *[]ProductImage   `json:"images" binding:"omitempty,dive"`
	CreatedAt     time.Time         `json:"createdAt" binding:"required"`
	UpdatedAt     time.Time         `json:"updatedAt" binding:"required"`
	DeletedAt     *time.Time        `json:"deletedAt" binding:"omitnil"`
	Category      *Category         `json:"category" binding:"omitnil"`
	Attributes    *[]Attribute      `json:"attributes" binding:"omitnil,dive"`
	Variants      *[]ProductVariant `json:"variants" binding:"omitnil,dive"`
}

type ProductOption struct {
	ID     int            `json:"id" binding:"required"`
	Name   string         `json:"name" binding:"required"`
	Values *[]OptionValue `json:"values" binding:"omitempty,dive"`
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
	DeletedAt     *time.Time      `json:"deletedAt" binding:"omitnil"`
	OptionValues  *[]OptionValue  `json:"optionValues" binding:"omitempty,dive"`
	Images        *[]ProductImage `json:"images" binding:"omitempty,dive"`
	Product       *Product        `json:"product,omitempty"`
}

type ProductImage struct {
	ID        int       `json:"id" binding:"required"`
	URL       string    `json:"url" binding:"required,url"`
	Order     int       `json:"order" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
}
