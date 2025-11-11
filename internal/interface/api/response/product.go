package response

import (
	"time"

	"backend/internal/domain/product"
)

type ProductOption struct {
	ID     int      `json:"id" binding:"required"`
	Name   string   `json:"name" binding:"required"`
	Values []string `json:"values" binding:"required"`
}

type ProductVariantOptionValue struct {
	ID    int    `json:"id" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type Product struct {
	ID              int               `json:"id" binding:"required"`
	Name            string            `json:"name" binding:"required"`
	Description     string            `json:"description" binding:"required"`
	ViewsCount      int               `json:"viewsCount" binding:"required"`
	TotalPurchase   int               `json:"totalPurchase" binding:"required"`
	TrendingScore   int64             `json:"trendingScore" binding:"required"`
	CreatedAt       time.Time         `json:"createdAt" binding:"required"`
	UpdatedAt       time.Time         `json:"updatedAt" binding:"required"`
	Category        Category          `json:"category" binding:"required"`
	AttributeValues []AttributeValue  `json:"attributeValues" binding:"required"`
	Variants        *[]ProductVariant `json:"variants" binding:"omitnil"`
}

// TODO: implement
// TODO: move to response?
func ProductFromDomain(p *product.Model) *Product {
	return &Product{}
}

func ProductsFromDomain(products *[]product.Model) []Product {
	return []Product{}
}

type ProductImage struct {
	ID               int       `json:"id" binding:"required"`
	URL              string    `json:"url" binding:"required"`
	Order            int       `json:"order" binding:"required"`
	CreatedAt        time.Time `json:"createdAt" binding:"required"`
	ProductVariantID *int      `json:"productVariantId" binding:"omitnil"`
}

type ProductVariantImage struct {
	ID        int       `json:"id" binding:"required"`
	URL       string    `json:"url" binding:"required"`
	Order     int       `json:"order" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
}

type ProductVariant struct {
	ID            int                         `json:"id" binding:"required"`
	SKU           string                      `json:"sku" binding:"required"`
	Price         int64                       `json:"price" binding:"required"`
	Quantity      int                         `json:"quantity" binding:"required"`
	PurchaseCount int                         `json:"purchaseCount" binding:"required"`
	CreatedAt     time.Time                   `json:"createdAt" binding:"required"`
	DeletedAt     *time.Time                  `json:"deletedAt" binding:"omitnil"`
	OptionValue   []ProductVariantOptionValue `json:"optionValues" binding:"required"`
	Images        []ProductVariantImage       `json:"images" binding:"omitnil"`
}

type ProductImageDeleteURL struct {
	URL string `json:"url" binding:"required"`
}

type ProductUploadURLImage struct {
	URL string `json:"url" binding:"required"`
	Key string `json:"key" binding:"required"`
}

// TODO: move to response?
func ProductUploadURLImageFromDomain(
	u *product.UploadImageURLModel,
) *ProductUploadURLImage {
	return &ProductUploadURLImage{
		URL: u.URL,
		Key: u.Key,
	}
}
