package http

import (
	"net/http"

	"backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandlerImpl struct {
	orderApp           application.Order
	ErrRequiredOrderID string
	ErrInvalidOrderID  string
}

var _ OrderHandler = &OrderHandlerImpl{}

func ProvideOrderHandler(orderApp application.Order) *OrderHandlerImpl {
	return &OrderHandlerImpl{
		orderApp:           orderApp,
		ErrRequiredOrderID: "order_id is required",
		ErrInvalidOrderID:  "invalid order_id",
	}
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
func (h *OrderHandlerImpl) Get(ctx *gin.Context) {
	orderIDString := ctx.Param("order_id")
	if orderIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredOrderID))
		return
	}
	orderID, err := uuid.Parse(orderIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidOrderID))
		return
	}
	order, err := h.orderApp.Get(ctx, application.GetOrderParam{
		OrderID: orderID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// ListOrders godoc
//
//	@Summary		List all orders
//	@Description	Get all orders
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_ids	query		[]string	false	"Filter by order IDs"	collectionFormat(csv)
//	@Param			user_ids	query		[]string	false	"Filter by user IDs"	collectionFormat(csv)
//	@Param			status_ids	query		[]string	false	"Filter by status IDs"	collectionFormat(csv)
//	@Success		200			{array}		domain.Order
//	@Failure		500			{object}	Error
//	@Router			/orders [get]
func (h *OrderHandlerImpl) List(ctx *gin.Context) {
	paginateParam, err := createPaginationParamsFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	var orderIDs *[]uuid.UUID
	if orderIDsQuery, ok := queryArrayToUUIDSlice(ctx, "order_ids"); ok {
		orderIDs = orderIDsQuery
	}

	var userIDs *[]uuid.UUID
	if userIDsQuery, ok := queryArrayToUUIDSlice(ctx, "user_ids"); ok {
		userIDs = userIDsQuery
	}

	var statusIDs *[]uuid.UUID
	if statusIDsQuery, ok := queryArrayToUUIDSlice(ctx, "status_ids"); ok {
		statusIDs = statusIDsQuery
	}

	orders, err := h.orderApp.List(ctx, application.ListOrderParam{
		PaginationParam: *paginateParam,
		IDs:             orderIDs,
		UserIDs:         userIDs,
		StatusIDs:       statusIDs,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, orders)
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
func (h *OrderHandlerImpl) Create(ctx *gin.Context) {
	var data application.CreateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	// TODO: Get userID from auth context
	userID := uuid.New() // Placeholder

	order, err := h.orderApp.Create(ctx, application.CreateOrderParam{
		UserID: userID,
		Data:   data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

// UpdateOrderStatus godoc
//
//	@Summary		Update order
//	@Description	Update  order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int							true	"Order ID"
//	@Param			status		body		application.UpdateOrderData	true	"Update order status request"
//	@Success		200			{object}	domain.Order
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{order_id} [patch]
func (h *OrderHandlerImpl) Update(ctx *gin.Context) {
	orderIDString := ctx.Param("order_id")
	if orderIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredOrderID))
		return
	}
	orderID, err := uuid.Parse(orderIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidOrderID))
		return
	}

	var data application.UpdateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	order, err := h.orderApp.Update(ctx, application.UpdateOrderParam{
		OrderID: orderID,
		Data:    data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, order)
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
func (h *OrderHandlerImpl) Delete(ctx *gin.Context) {
	orderIDString := ctx.Param("order_id")
	if orderIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredOrderID))
		return
	}
	orderID, err := uuid.Parse(orderIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidOrderID))
		return
	}

	err = h.orderApp.Delete(ctx, application.DeleteOrderParam{
		OrderID: orderID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
