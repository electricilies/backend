package response

import "backend/internal/domain/cart"

type CartItem struct {
	ID       int     `json:"id" binding:"required"`
	Product  Product `json:"product" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}

func CartItemFromDomain(i *cart.ItemModel) *CartItem {
	return &CartItem{
		ID:       i.ID,
		Product:  *ProductFromDomain(&i.Product),
		Quantity: i.Quantity,
	}
}

type Cart struct {
	ID    int                   `json:"id" binding:"required"`
	Items []CartItemsPagination `json:"items" binding:"required"`
}

func CartFromDomain(c *cart.Model) *Cart {
	cartItemsPagination := make([]CartItemsPagination, len(c.Items))
	for i, itemPagination := range c.Items {
		cartItemsPagination[i] = CartItemsPagination{
			Meta: *PaginationFromDomain(&itemPagination.Metadata),
			Data: CartItemsPaginationFromDomain(itemPagination.Items),
		}
	}

	return &Cart{
		ID:    c.ID,
		Items: cartItemsPagination,
	}
}

type CartItemsPagination struct {
	Meta Pagination `json:"meta" binding:"required"`
	Data []CartItem `json:"data" binding:"required"`
}

func CartItemsPaginationFromDomain(items []cart.ItemModel) []CartItem {
	cartItems := make([]CartItem, len(items))
	for i, item := range items {
		cartItems[i] = *CartItemFromDomain(&item)
	}
	return cartItems
}
