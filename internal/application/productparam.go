package application

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListProductParam struct {
	PaginationParam
	ProductIDs  *[]uuid.UUID        `binding:"omitempty"`
	CategoryIDs *[]uuid.UUID        `binding:"omitempty"`
	MinPrice    *int64              `binding:"omitempty"`
	MaxPrice    *int64              `binding:"omitempty"`
	Rating      *float64            `binding:"omitempty"`
	SortPrice   *string             `binding:"omitempty,oneof=asc desc"`
	SortRating  *string             `binding:"omitempty,oneof=asc desc"`
	Search      *string             `binding:"omitempty"`
	Deleted     domain.DeletedParam `binding:"omitempty,oneof=exclude only all"`
}

type CreateProductParam struct {
	Data CreateProductData `binding:"required"`
}

type CreateProductData struct {
	Name              string                     `json:"name"              binding:"required"`
	Description       string                     `json:"description"       binding:"required"`
	AttributeValueIDs *[]uuid.UUID               `json:"attributeValueIds" binding:"omitempty"`
	Options           []CreateProductOptionData  `json:"options"           binding:"required,dive"`
	CategoryID        uuid.UUID                  `json:"categoryId"        binding:"required"`
	Images            []CreateProductImageData   `json:"images"            binding:"required,dive"`
	Variants          []CreateProductVariantData `json:"variants"          binding:"required,dive"`
}

type CreateProductOptionData struct {
	Name string `json:"name" binding:"required"`
}

type CreateProductImageData struct {
	URL   string `json:"url"   binding:"required,url"`
	Order int    `json:"order" binding:"required"`
}

type CreateProductVariantData struct {
	SKU          string                             `json:"sku"          binding:"required"`
	Price        int64                              `json:"price"        binding:"required"`
	Quantity     int                                `json:"quantity"     binding:"required"`
	OptionValues *[]CreateProductVariantOptionValue `json:"optionValues" binding:"omitempty,dive"`
	Images       *[]CreateProductVariantImage       `json:"images"       binding:"omitempty,dive"`
}

type CreateProductVariantOptionValue struct {
	Name  string `json:"name"  binding:"required"`
	Value string `json:"value" binding:"required"`
}

type CreateProductVariantImage CreateProductImageData

type UpdateProductParam struct {
	ProductID uuid.UUID         `json:"productId" binding:"required"`
	Data      UpdateProductData `json:"data"      binding:"required"`
}

type UpdateProductData struct {
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	CategoryID  *uuid.UUID `json:"categoryId,omitempty"`
}

type GetProductParam struct {
	ProductID uuid.UUID `binding:"required"`
}

type DeleteProductParam struct {
	ProductID uuid.UUID `binding:"required"`
}

type AddProductImagesParam struct {
	Data []AddProductImageData `json:"data" binding:"required,dive"`
}

type AddProductImageData struct {
	URL              string     `json:"url"                        binding:"required"`
	Order            int        `json:"order,omitempty"`
	ProductID        uuid.UUID  `json:"productId,omitempty"`
	ProductVariantID *uuid.UUID `json:"productVariantId,omitempty"`
}

type DeleteProductImagesParam struct {
	IDs []uuid.UUID `json:"ids" binding:"required,dive"`
}

type AddProductVariantsParam struct {
	ProductID uuid.UUID                `json:"productId" binding:"required"`
	Data      []AddProductVariantsData `json:"data"      binding:"required,dive"`
}

type AddProductVariantsData struct {
	SKU            string      `json:"sku"                      binding:"required"`
	Price          int64       `json:"price"                    binding:"required"`
	Quantity       int         `json:"quantity"                 binding:"required"`
	OptionValueIDs []uuid.UUID `json:"optionValueIds,omitempty"`
}

type UpdateProductVariantParam struct {
	ProductID        uuid.UUID                `json:"productId"        binding:"required"`
	ProductVariantID uuid.UUID                `json:"ProductVariantId" binding:"required"`
	Data             UpdateProductVariantData `json:"data"             binding:"required"`
}

type UpdateProductVariantData struct {
	Price    *int64 `json:"price,omitempty"`
	Quantity *int   `json:"quantity,omitempty"`
}

type UpdateProductOptionsParam struct {
	ProductID uuid.UUID                  `json:"productId" binding:"required"`
	Data      []UpdateProductOptionsData `json:"data"      binding:"required,dive"`
}

type UpdateProductOptionsData struct {
	ID   uuid.UUID `json:"id"             binding:"required"`
	Name *string   `json:"name,omitempty"`
}

type UpdateProductOptionValuesParam struct {
	ProductID uuid.UUID                       `json:"productId" binding:"required"`
	Data      []UpdateProductOptionValuesData `json:"data"      binding:"required,dive"`
}

type UpdateProductOptionValuesData struct {
	ID    uuid.UUID `json:"id"              binding:"required"`
	Value *string   `json:"value,omitempty"`
}
