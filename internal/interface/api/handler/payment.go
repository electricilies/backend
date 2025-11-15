package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payment interface{}

type paymentHandler struct{}

func NewPayment() Payment { return &paymentHandler{} }

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
func (h *paymentHandler) Create(ctx *gin.Context) {
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
func (h *paymentHandler) Get(ctx *gin.Context) {
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
func (h *paymentHandler) List(ctx *gin.Context) {
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
func (h *paymentHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
