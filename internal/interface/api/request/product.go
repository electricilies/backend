package request

type CreateProduct struct {
	Name              string                 `json:"name" binding:"required"`
	Description       string                 `json:"description,omitempty"`
	CategoryIDs       []int                  `json:"categoryIds,omitempty"`
	AttributeValueIDs []int                  `json:"attributeValueIds,omitempty"`
	Category          int                    `json:"category" binding:"required"`
	ProductVariants   []CreateProductVariant `json:"productVariants" binding:"required"`
	ProductImages     []CreateProductImage   `json:"productImages" binding:"required"`
}

type UpdateProduct struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CategoryIds int    `json:"categoryIds,omitempty"`
}

type AddAttributeValues struct {
	AttributeValue []string `json:"attributeValue" binding:"required"`
}

type CreateProductVariant struct {
	SKU                 string `json:"sku" binding:"required"`
	Price               int64  `json:"price" binding:"required"`
	Quantity            int    `json:"quantity" binding:"required"`
	ProductOptionValues []int  `json:"productOptionValues,omitempty"`
}

type UpdateProductVariant struct {
	Price    int64 `json:"price,omitempty"`
	Quantity int   `json:"quantity,omitempty"`
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
	Value string `json:"value" binding:"required"`
}

type UpdateProductCategory struct {
	CategoryIDs []int `json:"categoryIds" binding:"required"`
}
