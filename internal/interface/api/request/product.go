package request

type productOption struct {
	Option      string `json:"option" binding:"required"`
	OptionValue string `json:"option_value" binding:"required"`
}

type CreateProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	CategoryIDs []int  `json:"category_ids,omitempty"`
}

type UpdateProduct struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type AddAttributeValues struct {
	AttributeValueIDs []int `json:"attribute_value_ids" binding:"required"`
}

type CreateProductVariant struct {
	SKU            string           `json:"sku" binding:"required"`
	Price          float64          `json:"price" binding:"required"`
	Quantity       int              `json:"quantity" binding:"required"`
	ProductOptions *[]productOption `json:"product_options,omitempty"`
}

type UpdateProductVariant struct {
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}

type CreateProductImage struct {
	URL              string `json:"url" binding:"required"`
	Order            int    `json:"order,omitempty"`
	ProductVariantID int    `json:"product_variant_id,omitempty"`
}
