package application

type GetCartParam struct {
	CartID int `binding:"required"`
}

type CreateCartParam struct {
	UserID int `binding:"required"`
}

type CreateCartItemParam struct {
	UserID int                `binding:"required"`
	CartID int                `binding:"required"`
	Data   CreateCartItemData `binding:"required"`
}

type CreateCartItemData struct {
	ProductVariantID int `json:"productVariantId" binding:"required"`
	Quantity         int `json:"quantity"         binding:"required"`
}

type UpdateCartItemParam struct {
	UserID int                `binding:"required"`
	CartID int                `binding:"required"`
	ItemID int                `binding:"required"`
	Data   UpdateCartItemData `binding:"required"`
}

type UpdateCartItemData struct {
	Quantity int `json:"quantity" binding:"required"`
}

type DeleteCartItemParam struct {
	UserID int `binding:"required"`
	CartID int `binding:"required"`
	ItemID int `binding:"required"`
}
