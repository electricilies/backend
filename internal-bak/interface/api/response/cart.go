package response

import "backend/internal/domain/cart"

type CartItem struct {
	ID             string         `json:"id" binding:"required"`
	ProductVariant ProductVariant `json:"productVariant" binding:"required"`
	Quantity       int            `json:"quantity" binding:"required"`
}

func CartItemFromDomain(i cart.ItemModel) *CartItem {
	return &CartItem{
		ID:             i.ID,
		ProductVariant: *ProductVariantFromDomain(i.ProductVariant),
		Quantity:       i.Quantity,
	}
}

type Cart struct {
	ID    int              `json:"id" binding:"required"`
	Items []DataPagination `json:"items" binding:"required"`
}

func CartFromDomain(c *cart.Model) *Cart {
	cartItemsPagination := make([]DataPagination, len(c.Items))
	for i, item := range c.Items {
		cartItemsPagination[i] = *DataPaginationFromDomain(item.Items, item.Metadata)
	}

	return &Cart{
		ID:    c.ID,
		Items: cartItemsPagination,
	}
}
