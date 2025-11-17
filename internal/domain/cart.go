package domain

type CartItem struct {
	ID             string         `json:"id" binding:"required"`
	ProductVariant ProductVariant `json:"productVariant" binding:"required"`
	Quantity       int            `json:"quantity" binding:"required"`
}

type Cart struct {
	ID    int              `json:"id" binding:"required"`
	Items []DataPagination `json:"items" binding:"required"`
}
