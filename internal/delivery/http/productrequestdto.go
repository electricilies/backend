package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListProductRequestDto struct {
	PaginationRequestDto
	ProductIDs  []uuid.UUID
	CategoryIDs []uuid.UUID
	MinPrice    int64
	MaxPrice    int64
	Rating      float64
	SortPrice   string
	SortRating  string
	Search      string
	Deleted     domain.DeletedParam
}

type CreateProductRequestDto struct {
	Data CreateProductData
}

type CreateProductData struct {
	Name              string                        `json:"name"                 binding:"required"`
	Description       string                        `json:"description"          binding:"required"`
	AttributeValueIDs []CreateProductAttributesData `json:"attributes,omitempty"`
	Options           []CreateProductOptionData     `json:"options,omitempty"    binding:"omitempty,dive"`
	CategoryID        uuid.UUID                     `json:"categoryId"           binding:"required"`
	Images            []CreateProductImageData      `json:"images"               binding:"required,dive"`
	Variants          []CreateProductVariantData    `json:"variants"             binding:"required,dive"`
}

type CreateProductAttributesData struct {
	AttributeID uuid.UUID `json:"attributeId" binding:"required"`
	ValueID     uuid.UUID `json:"valueId"     binding:"required"`
}

type CreateProductOptionData struct {
	Name   string   `json:"name"   binding:"required"`
	Values []string `json:"values" binding:"required"`
}

type CreateProductImageData struct {
	URL   string `json:"url"   binding:"required,url"`
	Order int    `json:"order" binding:"required"`
}

type CreateProductVariantData struct {
	SKU      string                       `json:"sku"               binding:"required"`
	Price    int64                        `json:"price"             binding:"required"`
	Quantity int                          `json:"quantity"          binding:"required"`
	Options  []CreateProductVariantOption `json:"options,omitempty" binding:"omitempty,dive"`
	Images   []CreateProductVariantImage  `json:"images,omitempty"  binding:"omitempty,dive"`
}

type CreateProductVariantOption struct {
	Name  string `json:"name"  binding:"required"`
	Value string `json:"value" binding:"required"`
}

type CreateProductVariantImage CreateProductImageData

type UpdateProductRequestDto struct {
	ProductID uuid.UUID
	Data      UpdateProductData
}

type UpdateProductData struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CategoryID  uuid.UUID `json:"categoryId,omitempty"`
}

type GetProductRequestDto struct {
	ProductID uuid.UUID
}

type DeleteProductRequestDto struct {
	ProductID uuid.UUID
}

type AddProductImagesRequestDto struct {
	ProductID uuid.UUID
	Data      []AddProductImageData
}

type AddProductImageData struct {
	URL              string    `json:"url"                        binding:"required"`
	Order            int       `json:"order,omitempty"`
	ProductVariantID uuid.UUID `json:"productVariantId,omitempty"`
}

type DeleteProductImagesRequestDto struct {
	ProductID uuid.UUID
	ImageIDs  []uuid.UUID
}

type AddProductVariantsRequestDto struct {
	ProductID uuid.UUID
	Data      []AddProductVariantsData
}

type AddProductVariantsData struct {
	SKU            string      `json:"sku"                      binding:"required"`
	Price          int64       `json:"price"                    binding:"required"`
	Quantity       int         `json:"quantity"                 binding:"required"`
	OptionValueIDs []uuid.UUID `json:"optionValueIds,omitempty"`
}

type UpdateProductVariantRequestDto struct {
	ProductID        uuid.UUID
	ProductVariantID uuid.UUID
	Data             UpdateProductVariantData
}

type UpdateProductVariantData struct {
	Price    int64 `json:"price,omitempty"`
	Quantity int   `json:"quantity,omitempty"`
}

type UpdateProductOptionsRequestDto struct {
	ProductID uuid.UUID
	Data      []UpdateProductOptionsData
}

type UpdateProductOptionsData struct {
	ID   uuid.UUID `json:"id"             binding:"required"`
	Name string    `json:"name,omitempty"`
}

type UpdateProductOptionValuesRequestDto struct {
	ProductID uuid.UUID
	OptionID  uuid.UUID
	Data      []UpdateProductOptionValuesData
}

type UpdateProductOptionValuesData struct {
	ID    uuid.UUID `json:"id"              binding:"required"`
	Value string    `json:"value,omitempty"`
}
