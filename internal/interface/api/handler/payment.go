package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payment interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type PaymentHandler struct{}

func NewPayment() Payment { return &PaymentHandler{} }

func ProvidePayment() *PaymentHandler {
	return &PaymentHandler{}
}

// CreatePayment godoc
//
//	@Summary		Create a new payment
//	@Description	Create a new payment
//	@Tags			Payment
//	@Accept			json
//	@Produce		json
//	@Param			payment	body		request.CreatePayment	true	"Payment request"
//	@Success		201		{object}	response.Payment
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/payments [post]
func (h *PaymentHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// GetPayment godoc
//
//	@Summary		Get payment by ID
//	@Description	Get payment details by ID
//	@Tags			Payment
//	@Accept			json
//	@Produce		json
//	@Param			payment_id	path		int	true	"Payment ID"
//	@Success		200			{object}	response.Payment
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/payments/{payment_id} [get]
func (h *PaymentHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListPayments godoc
//
//	@Summary		List all payments
//	@Description	Get all payments
//	@Tags			Payment
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Payment
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/payments [get]
func (h *PaymentHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdatePayment godoc
//
//	@Summary		Update payment by ID
//	@Description	Update payment details by ID
//	@Tags			Payment
//	@Accept			json
//	@Produce		json
//	@Param			payment_id	path		int						true	"Payment ID"
//	@Param			payment		body		request.UpdatePayment	true	"Payment update request"
//	@Success		200			{object}	response.Payment
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/payments/{payment_id} [put]
func (h *PaymentHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
