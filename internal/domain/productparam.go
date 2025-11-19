package domain

type CreateProductParam struct {
	Data CreateProductData `binding:"required"`
}

type CreateProductData struct {
	Name              string                     `json:"name" binding:"required"`
	Description       string                     `json:"description" binding:"required"`
	AttributeValueIDs *[]int                     `json:"attributeValueIds" binding:"omitempty"`
	Options           []CreateProductOptionData  `json:"options" binding:"required,dive"`
	Category          int                        `json:"category" binding:"required"`
	Images            []CreateProductImageData   `json:"images" binding:"required,dive"`
	Variants          []CreateProductVariantData `json:"variants" binding:"required,dive"`
}

type CreateProductOptionData struct {
	Name string `json:"name" binding:"required"`
}

type CreateProductImageData struct {
	URL   string `json:"url" binding:"required,url"`
	Order int    `json:"order,omitempty"`
}

type CreateProductVariantData struct {
	SKU          string                             `json:"sku" binding:"required"`
	Price        int64                              `json:"price" binding:"required"`
	Quantity     int                                `json:"quantity" binding:"required"`
	OptionValues *[]CreateProductVariantOptionValue `json:"optionValues" binding:"omitempty,dive"`
	Images       *[]CreateProductVariantImage       `json:"images" binding:"omitempty,dive"`
}

type CreateProductVariantOptionValue struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type CreateProductVariantImage CreateProductImageData

// Update Product

type UpdateProductParam struct {
	ProductID int               `json:"productId" binding:"required"`
	Data      UpdateProductData `json:"data" binding:"required"`
}

type UpdateProductData struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryID  *int    `json:"categoryId,omitempty"`
}

// Get

type GetProductParam struct {
	ProductID int `binding:"required"`
}

// Delete

type DeleteProductParam struct {
	ProductID int `binding:"required"`
}

// Update option value

type UpdateProductOptionParam struct {
	OptionID int     `json:"id" binding:"required"`
	Name     *string `json:"name,omitempty"`
}

type CreateProductVariantParam struct {
	SKU                 string   `json:"sku" binding:"required"`
	Price               int64    `json:"price" binding:"required"`
	Quantity            int      `json:"quantity" binding:"required"`
	ProductOptionValues []string `json:"productOptionValues,omitempty"`
}

type UpdateProductVariantParam struct {
	Price    *int64 `json:"price,omitempty"`
	Quantity *int   `json:"quantity,omitempty"`
}

type CreateProductImageParam struct {
	URL              string `json:"url" binding:"required"`
	Order            int    `json:"order,omitempty"`
	ProductVariantID int    `json:"productVariantId,omitempty"`
	ProductID        int    `json:"productId,omitempty"`
}
