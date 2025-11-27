package http

import "github.com/google/uuid"

type GetCartRequestDto struct {
	CartID uuid.UUID
}

type GetCartByUserRequestDto struct {
	UserID uuid.UUID
}

type CreateCartRequestDto struct {
	Data CreateCartData
}

type CreateCartData struct {
	UserID uuid.UUID `json:"userId" binding:"required"`
}

type CreateCartItemRequestDto struct {
	UserID uuid.UUID
	CartID uuid.UUID
	Data   CreateCartItemData
}

type CreateCartItemData struct {
	ProductID        uuid.UUID `json:"productId"        binding:"required"`
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required,gt=0"`
}

type UpdateCartItemRequestDto struct {
	UserID uuid.UUID
	CartID uuid.UUID
	ItemID uuid.UUID
	Data   UpdateCartItemData
}

type UpdateCartItemData struct {
	Quantity int `json:"quantity"`
}

type DeleteCartItemRequestDto struct {
	UserID uuid.UUID
	CartID uuid.UUID
	ItemID uuid.UUID
}
