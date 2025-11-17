package request

import "backend/internal/domain/product"

type CreateProduct struct {
	Name              string                 `json:"name" binding:"required"`
	Description       string                 `json:"description,omitempty"`
	CategoryIDs       []int                  `json:"categoryIds,omitempty"`
	AttributeValueIDs []int                  `json:"attributeValueIds,omitempty"`
	ProductOption     []CreateProductOption  `json:"productOption,omitempty"`
	Category          int                    `json:"category" binding:"required"`
	ProductVariants   []CreateProductVariant `json:"productVariants" binding:"required"`
	ProductImages     []CreateProductImage   `json:"productImages" binding:"required"`
}

type UpdateProduct struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryID  *int    `json:"categoryId,omitempty"`
}

type AddAttributeValues struct {
	AttributeValue []string `json:"attributeValue" binding:"required"`
}

type CreateProductVariant struct {
	SKU                 string   `json:"sku" binding:"required"`
	Price               int64    `json:"price" binding:"required"`
	Quantity            int      `json:"quantity" binding:"required"`
	ProductOptionValues []string `json:"productOptionValues,omitempty"`
}

type UpdateProductVariant struct {
	Price    *int64 `json:"price,omitempty"`
	Quantity *int   `json:"quantity,omitempty"`
}

type CreateProductImage struct {
	URL              string `json:"url" binding:"required"`
	Order            int    `json:"order,omitempty"`
	ProductVariantID int    `json:"productVariantId,omitempty"`
	ProductID        int    `json:"productId,omitempty"`
}

type CreateProductOption struct {
	Option string   `json:"option" binding:"required"`
	Value  []string `json:"value" binding:"required"`
}

type UpdateProductOption struct {
	Name *string `json:"name" binding:"required"`
}

type ProductQueryParams struct {
	Limit       int
	Offset      int
	CategoryIDs []int
	MinPrice    int64
	MaxPrice    int64
	SortPrice   string
	SortRating  string
	Search      string
	Deleted     string
}

func ProductQueryParamsToDomain(productQueryParams ProductQueryParams) *product.QueryParams {
	return &product.QueryParams{
		PaginationParams: *PaginationParamsToDomain(productQueryParams.Limit, productQueryParams.Offset),
		Search:           &productQueryParams.Search,
		MinPrice:         &productQueryParams.MinPrice,
		MaxPrice:         &productQueryParams.MaxPrice,
		CategoryIDs:      &productQueryParams.CategoryIDs,
		Deleted:          *DeletedParamToDomain(productQueryParams.Deleted),
		SortPrice:        *SortPriceParamToDomain(productQueryParams.SortPrice),
		SortRating:       *SortRatingParamToDomain(productQueryParams.SortRating),
	}
}
