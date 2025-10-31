package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ListItemByOrder(ctx *gin.Context)
}

type orderHandler struct{}

func NewOrder() Order { return &orderHandler{} }

// GetOrder godoc
//
//	@Summary		Get order by ID
//	@Description	Get order details by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200	{object}	response.Order
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/orders/{id} [get]
func (h *orderHandler) Get(ctx *gin.Context) {
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
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/orders [get]
func (h *orderHandler) List(ctx *gin.Context) {
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
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		409		{object}	mapper.ConflictError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/orders [post]
func (h *orderHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateOrderStatus godoc
//
//	@Summary		Update order status
//	@Description	Update the status of an order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Order ID"
//	@Param			status	body		request.UpdateOrderStatus	true	"Update order status request"
//	@Success		204		{string}	string						"no content"
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		404		{object}	mapper.NotFoundError
//	@Failure		409		{object}	mapper.ConflictError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/orders/{id}/status [put]
func (h *orderHandler) UpdateStatus(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteOrder godoc
//
//	@Summary		Delete order
//	@Description	Delete an order by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Order ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/orders/{id} [delete]
func (h *orderHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListOrderItems godoc
//
//	@Summary		List items by order
//	@Description	Get all items for a specific order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200	{array}		response.OrderItem
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/orders/{id}/items [get]
func (h *orderHandler) ListItemByOrder(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
