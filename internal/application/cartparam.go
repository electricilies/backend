package application

import "github.com/google/uuid"

type GetCartParam struct {
	CartID uuid.UUID `binding:"required"`
}

type CreateCartParam struct {
	UserID uuid.UUID `binding:"required"`
}

type CreateCartItemParam struct {
	UserID uuid.UUID          `binding:"required"`
	CartID uuid.UUID          `binding:"required"`
	Data   CreateCartItemData `binding:"required"`
}

type CreateCartItemData struct {
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required"`
}

type UpdateCartItemParam struct {
	UserID uuid.UUID          `binding:"required"`
	CartID uuid.UUID          `binding:"required"`
	ItemID uuid.UUID          `binding:"required"`
	Data   UpdateCartItemData `binding:"required"`
}

type UpdateCartItemData struct {
	Quantity int `json:"quantity" binding:"required"`
}

type DeleteCartItemParam struct {
	UserID uuid.UUID `binding:"required"`
	CartID uuid.UUID `binding:"required"`
	ItemID uuid.UUID `binding:"required"`
}
