package handler

import (
	"net/http"
	"strconv"

	"backend/internal/application"
	"backend/internal/interface/api/request"
	"backend/internal/interface/api/response"

	"github.com/gin-gonic/gin"
)

type Cart interface {
	Get(ctx *gin.Context)
	CreateItem(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	RemoveItem(ctx *gin.Context)
}

type CartHandler struct {
	app application.Cart
}

func NewCart() Cart { return &CartHandler{} }

func ProvideCart(
	app application.Cart,
) *CartHandler {
	return &CartHandler{
		app: app,
	}
}

// GetCart godoc
//
//	@Summary		Get cart
//	@Description	Get cart by user ID
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"Limit"	default(20)
//	@Param			offset	query		int	false	"Offset"
//	@Success		200		{object}	response.Cart{items=response.DataPagination{data=[]response.CartItem}}
//	@Failure		404		{object}	response.NotFoundError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/carts/{cart_id} [get]
func (h *CartHandler) Get(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	id, _ := strconv.Atoi(ctx.Param("id"))

	cart, err := h.app.Get(
		ctx,
		id,
		request.CartQueryParamsToDomain(&request.CartQueryParams{
			Limit:  limit,
			Offset: offset,
		}),
	)
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.CartFromDomain(cart))
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
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/carts/{cart_id}/item  [post]
func (h *CartHandler) CreateItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateCartItem godoc
//
//	@Summary		Update cart item
//	@Description	Update quantity of a cart item
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart_item_id	path		int						true	"Cart Item ID"
//	@Param			item			body		request.UpdateCartItem	true	"Update cart item request"
//	@Success		200				{object}	response.CartItem
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		404				{object}	response.NotFoundError
//	@Failure		500				{object}	response.InternalServerError
//
// /	@Router			/carts/{cart_id}/item [patch]
func (h *CartHandler) UpdateItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// RemoveCartItem godoc
//
//	@Summary		Remove cart item
//	@Description	Remove an item from the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			item_id	path	int	true	"Cart Item ID"
//	@Success		204
//	@Failure		404	{object}	response.NotFoundError
//	@Failure		500	{object}	response.InternalServerError
//
// /	@Router			/carts/{cart_id}/item [delete]
func (h *CartHandler) RemoveItem(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
