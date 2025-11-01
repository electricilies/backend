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
    ViewsCount      int              `json:"viewsCount" binding:"required"`
    TotalPurchase   int              `json:"totalPurchase" binding:"required"`
    TrendingScore   float64          `json:"trendingScore" binding:"required"`
    CreatedAt       time.Time        `json:"createdAt" binding:"required"`
    UpdatedAt       time.Time        `json:"updatedAt" binding:"required"`
    Categories      []string         `json:"categories" binding:"required"`
    AttributeValues []AttributeValue `json:"attributeValues" binding:"required"`
    Variants        []ProductVariant `json:"variants" binding:"required"`
}

type ProductVariantImage struct {
    ID        int       `json:"id" binding:"required"`
    URL       string    `json:"url" binding:"required"`
    Order     int       `json:"order" binding:"required"`
    CreatedAt time.Time `json:"createdAt" binding:"required"`
}

type ProductVariant struct {
    ID                 int                   `json:"id" binding:"required"`
    SKU                string                `json:"sku" binding:"required"`
    Price              float64               `json:"price" binding:"required"`
    Quantity           int                   `json:"quantity" binding:"required"`
    PurchaseCount      int                   `json:"purchaseCount" binding:"required"`
    CreatedAt          time.Time             `json:"createdAt" binding:"required"`
    DeletedAt          *time.Time            `json:"deletedAt" binding:"omitnil"`
    ProductOptionValue []ProductOptionValue  `json:"productOptions" binding:"required"`
    Images             []ProductVariantImage `json:"images" binding:"required"`
}
