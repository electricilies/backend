package service

type GetCartParam struct {
	CartID int `json:"cartId" binding:"required"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type AddCartItemParam struct {
	CartID    int `json:"cartId" binding:"required"`
	ProductID int `json:"productId" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type UpdateCartItemParam struct {
	CartID   int `json:"cartId" binding:"required"`
	ItemID   int `json:"itemId" binding:"required,uuid"`
	Quantity int `json:"quantity" binding:"required"`
}

type RemoveCartItemParam struct {
	CartID int `json:"cartId" binding:"required"`
	ItemID int `json:"itemId" binding:"required,uuid"`
}
