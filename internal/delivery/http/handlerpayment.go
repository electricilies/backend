package http

import (
	"net/http"

	_ "backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type PaymentHandler interface {
	Create(*gin.Context)
	Get(*gin.Context)
	List(*gin.Context)
	Update(*gin.Context)
}

type GinPaymentHandler struct{}

var _ PaymentHandler = &GinPaymentHandler{}

func ProvidePaymentHandler() *GinPaymentHandler {
	return &GinPaymentHandler{}
}

// CreatePayment godoc
//
//	@Summary		Create a new payment
//	@Description	Create a new payment
//	@Tags			Payment
//	@Accept			json
//	@Produce		json
//	@Param			payment	body		domain.CreatePaymentData	true	"Payment request"
//	@Success		201		{object} domain.Payment
//	@Failure		400		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/payments [post]
func (h *GinPaymentHandler) Create(ctx *gin.Context) {
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
//	@Success		200			{object} domain.Payment
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/payments/{payment_id} [get]
func (h *GinPaymentHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListPayments godoc
//
//	@Summary		List all payments
//	@Description	Get all payments
//	@Tags			Payment
//	@Accept			json
//	@Produce		json
//	@Success		200	{array} domain.Payment
//	@Failure		500	{object}	Error
//	@Router			/payments [get]
func (h *GinPaymentHandler) List(ctx *gin.Context) {
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
//	@Param			payment		body		domain.UpdatePaymentData	true	"Payment update request"
//	@Success		200			{object} domain.Payment
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/payments/{payment_id} [put]
func (h *GinPaymentHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
