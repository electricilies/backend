package http

import "github.com/google/uuid"

type GetCartRequestDto struct {
	CartID uuid.UUID `binding:"required"`
}

type CreateCartRequestDto struct {
	UserID uuid.UUID `binding:"required"`
}

type CreateCartItemRequestDto struct {
	UserID uuid.UUID          `binding:"required"`
	CartID uuid.UUID          `binding:"required"`
	Data   CreateCartItemData `binding:"required"`
}

type CreateCartItemData struct {
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required"`
}

type UpdateCartItemRequestDto struct {
	UserID uuid.UUID          `binding:"required"`
	CartID uuid.UUID          `binding:"required"`
	ItemID uuid.UUID          `binding:"required"`
	Data   UpdateCartItemData `binding:"required"`
}

type UpdateCartItemData struct {
	Quantity int `json:"quantity" binding:"required"`
}

type DeleteCartItemRequestDto struct {
	UserID uuid.UUID `binding:"required"`
	CartID uuid.UUID `binding:"required"`
	ItemID uuid.UUID `binding:"required"`
}
