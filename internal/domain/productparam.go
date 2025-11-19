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

type UpdateProductParam struct {
	ProductID int               `json:"productId" binding:"required"`
	Data      UpdateProductData `json:"data" binding:"required"`
}

type UpdateProductData struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryID  *int    `json:"categoryId,omitempty"`
}

type GetProductParam struct {
	ProductID int `binding:"required"`
}

type DeleteProductParam struct {
	ProductID int `binding:"required"`
}

type AddProductImagesParam struct {
	Data []AddProductImageData `json:"data" binding:"required,dive"`
}

type AddProductImageData struct {
	URL              string `json:"url" binding:"required"`
	Order            int    `json:"order,omitempty"`
	ProductID        int    `json:"productId,omitempty"`
	ProductVariantID *int   `json:"productVariantId,omitempty"`
}

type DeleteProductImagesParam struct {
	IDs []int `json:"ids" binding:"required,dive"`
}

type AddProductVariantsParam struct {
	ProductID int                      `json:"productId" binding:"required"`
	Data      []AddProductVariantsData `json:"data" binding:"required,dive"`
}

type AddProductVariantsData struct {
	SKU            string `json:"sku" binding:"required"`
	Price          int64  `json:"price" binding:"required"`
	Quantity       int    `json:"quantity" binding:"required"`
	OptionValueIDs []int  `json:"optionValueIds,omitempty"`
}

type UpdateProductVariantParam struct {
	ID   int                      `json:"id" binding:"required"`
	Data UpdateProductVariantData `json:"data" binding:"required"`
}

type UpdateProductVariantData struct {
	Price    *int64 `json:"price,omitempty"`
	Quantity *int   `json:"quantity,omitempty"`
}

type UpdateProductOptionsParam struct {
	Data []UpdateProductOptionsData `json:"data" binding:"required,dive"`
}

type UpdateProductOptionsData struct {
	ID   int     `json:"id" binding:"required"`
	Name *string `json:"name,omitempty"`
}

type UpdateProductOptionValuesParam struct {
	Data []UpdateProductOptionValuesData `json:"data" binding:"required,dive"`
}

type UpdateProductOptionValuesData struct {
	ID    int     `json:"id" binding:"required"`
	Value *string `json:"value,omitempty"`
}
