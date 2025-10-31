package response

import "time"

type ProductOptionValue struct {
	ID    int    `json:"id" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type Product struct {
	ID              int              `json:"id" binding:"required"`
	Name            string           `json:"name" binding:"required"`
	Description     string           `json:"description" binding:"required"`
	ViewsCount      int              `json:"views_count" binding:"required"`
	TotalPurchase   int              `json:"total_purchase" binding:"required"`
	TrendingScore   float64          `json:"trending_score" binding:"required"`
	CreatedAt       time.Time        `json:"created_at" binding:"required"`
	UpdatedAt       time.Time        `json:"updated_at" binding:"required"`
	Categories      []string         `json:"categories" binding:"required"`
	AttributeValues []AttributeValue `json:"attribute_values" binding:"required"`
	Variants        []ProductVariant `json:"variants" binding:"required"`
}

type ProductVariantImage struct {
	ID        int       `json:"id" binding:"required"`
	URL       string    `json:"url" binding:"required"`
	Order     int       `json:"order" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

type ProductVariant struct {
	ID                 int                   `json:"id" binding:"required"`
	SKU                string                `json:"sku" binding:"required"`
	Price              float64               `json:"price" binding:"required"`
	Quantity           int                   `json:"quantity" binding:"required"`
	PurchaseCount      int                   `json:"purchase_count" binding:"required"`
	CreatedAt          time.Time             `json:"created_at" binding:"required"`
	DeletedAt          *time.Time            `json:"deleted_at" binding:"omitnil"`
	ProductOptionValue []ProductOptionValue  `json:"product_options" binding:"required"`
	Images             []ProductVariantImage `json:"images," binding:"required"`
}
