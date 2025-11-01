package request

type AddCartItem struct {
	ProductID        int `json:"productId" binding:"required"`
	ProductVariantID int `json:"productVariantId" binding:"required"`
	Quantity         int `json:"quantity" binding:"required"`
}

type UpdateCartItem struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}
