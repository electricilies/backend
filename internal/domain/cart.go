package domain

type Cart struct {
	ID    int        `json:"id" binding:"required"`
	Items []CartItem `json:"items" binding:"required"`
}

type CartItem struct {
	ID             string         `json:"id" binding:"required"`
	ProductVariant ProductVariant `json:"productVariant" binding:"required"`
	Quantity       int            `json:"quantity" binding:"required"`
}
