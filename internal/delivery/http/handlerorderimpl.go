package http

import (
	"net/http"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandlerImpl struct {
	orderApp           OrderApplication
	ErrRequiredOrderID string
	ErrInvalidOrderID  string
}

var _ OrderHandler = &OrderHandlerImpl{}

func ProvideOrderHandler(orderApp OrderApplication) *OrderHandlerImpl {
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
//	@Param			order_id	path		string	true	"Order ID"	format(uuid)
//	@Success		200			{object}	OrderResponseDto
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{order_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
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
	order, err := h.orderApp.Get(ctx, GetOrderRequestDto{
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
//	@Param			order_ids	query		[]string			false	"Filter by order IDs"	collectionFormat(csv)	format(uuid)
//	@Param			user_ids	query		[]string			false	"Filter by user IDs"	collectionFormat(csv)	format(uuid)
//	@Param			status		query		domain.OrderStatus	false	"Filter by statuses"
//	@Success		200			{array}		OrderResponseDto
//	@Failure		500			{object}	Error
//	@Router			/orders [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *OrderHandlerImpl) List(ctx *gin.Context) {
	paginateRequestDto, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	orderIDs, _ := queryArrayToUUIDSlice(ctx, "order_ids")
	userIDs, _ := queryArrayToUUIDSlice(ctx, "user_ids")

	statusQuery := ctx.Query("status")
	status := domain.OrderStatus(statusQuery)

	orders, err := h.orderApp.List(ctx, ListOrderRequestDto{
		PaginationRequestDto: *paginateRequestDto,
		IDs:                  orderIDs,
		UserIDs:              userIDs,
		Status:               status,
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
//	@Param			order	body		CreateOrderData	true	"Order request"
//	@Success		201		{object}	OrderResponseDto
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/orders [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *OrderHandlerImpl) Create(ctx *gin.Context) {
	var data CreateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	order, err := h.orderApp.Create(ctx, CreateOrderRequestDto{
		Data: data,
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
//	@Param			order_id	path		string			true	"Order ID"	format(uuid)
//	@Param			status		body		UpdateOrderData	true	"Update order status request"
//	@Success		200			{object}	OrderResponseDto
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{order_id} [put]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
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

	var data UpdateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	order, err := h.orderApp.Update(ctx, UpdateOrderRequestDto{
		OrderID: orderID,
		Data:    data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// VerifyVNPayIPN godoc
//
//	@Summary		Verify VNPay IPN
//	@Description	Verify VNPay IPN
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			param	query		VerifyVNPayIPNQueryParams	true	"VNPay IPN data"
//	@Success		200		{object}	VerifyVNPayIPNResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/orders/vnpay/ipn [get]
func (h *OrderHandlerImpl) VerifyVNPayIPN(ctx *gin.Context) {
	var queryParams VerifyVNPayIPNQueryParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	response, err := h.orderApp.VerifyVNPayIPN(ctx.Request.Context(), VerifyVNPayIPNRequestDTO{
		QueryParams: &queryParams,
	})
	if err != nil {
		SendVNPayError(ctx, response, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
