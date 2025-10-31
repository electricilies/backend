package response

import "time"

type ProductOptionValue struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Product struct {
	ID              int               `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	ViewsCount      int               `json:"views_count"`
	TotalPurchase   int               `json:"total_purchase"`
	TrendingScore   float64           `json:"trending_score"`
	CreatedAt       *time.Time        `json:"created_at"`
	UpdatedAt       *time.Time        `json:"updated_at"`
	Categories      []string          `json:"categories,omitempty"`
	AttributeValues *[]AttributeValue `json:"attribute_values,omitempty"`
	Variants        *[]ProductVariant `json:"variants,omitempty"`
	Images          *[]ProductImage   `json:"images,omitempty"`
}

type ProductImage struct {
	ID               int        `json:"id"`
	URL              string     `json:"url"`
	Order            int        `json:"order"`
	ProductVariantID *int       `json:"product_variant_id,omitempty"`
	CreatedAt        *time.Time `json:"created_at"`
}

type ProductVariant struct {
	ID                 int                   `json:"id"`
	SKU                string                `json:"sku"`
	Price              float64               `json:"price"`
	Quantity           int                   `json:"quantity"`
	PurchaseCount      int                   `json:"purchase_count"`
	CreatedAt          *time.Time            `json:"created_at"`
	DeletedAt          *time.Time            `json:"deleted_at,omitempty"`
	ProductOptionValue *[]ProductOptionValue `json:"product_options,omitempty"`
}
