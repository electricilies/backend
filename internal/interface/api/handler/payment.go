package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payment interface {
	ListMethods(ctx *gin.Context)
}

type paymentHandler struct{}

func NewPayment() Payment { return &paymentHandler{} }

// ListPaymentMethods godoc
//
// @Summary      List payment methods
// @Description  Get all available payment methods
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Success      200 {array} response.PaymentMethod
// @Failure      500 {object} mapper.InternalServerError
// @Router       /payments/methods [get]
func (h *paymentHandler) ListMethods(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
