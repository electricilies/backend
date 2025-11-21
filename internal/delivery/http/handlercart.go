package http

import (
	"net/http"

	_ "backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	Get(*gin.Context)
	CreateItem(*gin.Context)
	UpdateItem(*gin.Context)
	RemoveItem(*gin.Context)
}

type GinCartHandler struct{}

var _ CartHandler = &GinCartHandler{}

func ProvideCartHandler() *GinCartHandler {
	return &GinCartHandler{}
}

// GetCart godoc
//
//	@Summary		Get cart
//	@Description	Get cart by user ID
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"Limit"	default(20)
//	@Param			page	query		int	false	"Page"
//	@Success		200		{object}	domain.Cart
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id} [get]
func (h *GinCartHandler) Get(ctx *gin.Context) {}

// AddCartItem godoc
//
//	@Summary		Add item to cart
//	@Description	Add a product item to the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			item	body		application.CreateCartItemData	true	"Cart item request"
//	@Success		201		{object}	domain.CartItem
//	@Failure		400		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id}/item  [post]
func (h *GinCartHandler) CreateItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateCartItem godoc
//
//		@Summary		Update cart item
//		@Description	Update quantity of a cart item
//		@Tags			Cart
//		@Accept			json
//		@Produce		json
//		@Param			cart_item_id	path		int							true	"Cart Item ID"
//		@Param			item			body		application.UpdateCartItemData	true	"Update cart item request"
//		@Success		200				{object}	domain.CartItem
//		@Failure		400				{object}	Error
//		@Failure		404				{object}	Error
//		@Failure		500				{object}	Error
//	 	@Router			/carts/{cart_id}/item [patch]
func (h *GinCartHandler) UpdateItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// RemoveCartItem godoc
//
//		@Summary		Remove cart item
//		@Description	Remove an item from the cart
//		@Tags			Cart
//		@Accept			json
//		@Produce		json
//		@Param			item_id	path	int	true	"Cart Item ID"
//		@Success		204
//		@Failure		404	{object}	Error
//		@Failure		500	{object}	Error
//	 	@Router			/carts/{cart_id}/item [delete]
func (h *GinCartHandler) RemoveItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
