package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type OrderImpl struct{}

func ProvideOrder() *OrderImpl {
	return &OrderImpl{}
}

func NewOrder() Order { return &OrderImpl{} }

// GetOrder godoc
//
//	@Summary		Get order by ID
//	@Description	Get order details by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int	true	"Order ID"
//	@Success		200			{object}	response.Order
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/orders/{order_id} [get]
func (h *OrderImpl) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListOrders godoc
//
//	@Summary		List all orders
//	@Description	Get all orders
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Order
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/orders [get]
func (h *OrderImpl) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Create a new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order	body		request.CreateOrder	true	"Order request"
//	@Success		201		{object}	response.Order
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		409		{object}	response.ConflictError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/orders [post]
func (h *OrderImpl) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateOrderStatus godoc
//
//	@Summary		Update order
//	@Description	Update  order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int							true	"Order ID"
//	@Param			status		body		request.UpdateOrderStatus	true	"Update order status request"
//	@Success		200			{object}	response.Order
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		409			{object}	response.ConflictError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/orders/{order_id} [patch]
func (h *OrderImpl) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteOrder godoc
//
//	@Summary		Delete order
//	@Description	Delete an order by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int		true	"Order ID"
//	@Success		204			{string}	string	"no content"
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/orders/{order_id} [delete]
func (h *OrderImpl) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
