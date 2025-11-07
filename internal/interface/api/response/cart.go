package response

type CartItem struct {
	ID       int                 `json:"id" binding:"required"`
	Product  ProductWithVariants `json:"product" binding:"required"`
	Quantity int                 `json:"quantity" binding:"required"`
}

type Cart struct {
	ID    int        `json:"id" binding:"required"`
	Items []CartItem `json:"items" binding:"required"`
}
