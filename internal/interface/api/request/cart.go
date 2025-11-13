package request

import "backend/internal/domain/cart"

type AddCartItem struct {
	ProductID        int `json:"productId" binding:"required"`
	ProductVariantID int `json:"productVariantId" binding:"required"`
	Quantity         int `json:"quantity" binding:"required"`
}

type UpdateCartItem struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

type CartQueryParams struct {
	Limit  int
	Offset int
}

func CartQueryParamsToDomain(cartQueryParams *CartQueryParams) *cart.QueryParams {
	return &cart.QueryParams{
		Pagination: PaginationParamsToDomain(
			cartQueryParams.Limit,
			cartQueryParams.Offset,
		),
	}
}
