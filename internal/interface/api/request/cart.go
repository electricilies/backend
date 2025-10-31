package request

type AddCartItem struct {
	ProductID        int `json:"product_id" binding:"required"`
	ProductVariantID int `json:"product_variant_id" binding:"required"`
	Quantity         int `json:"quantity" binding:"required"`
}

type UpdateCartItem struct {
	Quantity int `json:"quantity" binding:"required"`
}
