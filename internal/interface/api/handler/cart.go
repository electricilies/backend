package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cart interface {
	Get(ctx *gin.Context)
	AddItem(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	RemoveItem(ctx *gin.Context)
}

type cartHandler struct{}

func NewCart() Cart { return &cartHandler{} }

// GetCart godoc
//
//	@Summary		Get cart
//	@Description	Get cart for the current user
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Cart
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/cart [get]
func (h *cartHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// AddCartItem godoc
//
//	@Summary		Add item to cart
//	@Description	Add a product item to the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			item	body		request.AddCartItem	true	"Cart item request"
//	@Success		201		{object}	response.CartItem
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/cart/items [post]
func (h *cartHandler) AddItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateCartItem godoc
//
//	@Summary		Update cart item
//	@Description	Update quantity of a cart item
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Cart Item ID"
//	@Param			item	body		request.UpdateCartItem	true	"Update cart item request"
//	@Success		204		{string}	string					"no content"
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		404		{object}	mapper.NotFoundError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/cart/items/{id} [put]
func (h *cartHandler) UpdateItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// RemoveCartItem godoc
//
//	@Summary		Remove cart item
//	@Description	Remove an item from the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Cart Item ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/cart/items/{id} [delete]
func (h *cartHandler) RemoveItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
