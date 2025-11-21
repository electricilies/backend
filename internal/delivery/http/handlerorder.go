package http

import (
	"net/http"

	_ "backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type GinOrderHandler struct{}

var _ OrderHandler = &GinOrderHandler{}

func ProvideOrderHandler() *GinOrderHandler {
	return &GinOrderHandler{}
}

// GetOrder godoc
//
//	@Summary		Get order by ID
//	@Description	Get order details by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int	true	"Order ID"
//	@Success		200			{object}	domain.Order
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{order_id} [get]
func (h *GinOrderHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListOrders godoc
//
//	@Summary		List all orders
//	@Description	Get all orders
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Order
//	@Failure		500	{object}	Error
//	@Router			/orders [get]
func (h *GinOrderHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Create a new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order	body		application.CreateOrderData	true	"Order request"
//	@Success		201		{object}	domain.Order
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/orders [post]
func (h *GinOrderHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateOrderStatus godoc
//
//	@Summary		Update order
//	@Description	Update  order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int						true	"Order ID"
//	@Param			status		body		application.UpdateOrderData	true	"Update order status request"
//	@Success		200			{object}	domain.Order
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{order_id} [patch]
func (h *GinOrderHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteOrder godoc
//
//	@Summary		Delete order
//	@Description	Delete an order by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path	int	true	"Order ID"
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/orders/{order_id} [delete]
func (h *GinOrderHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
