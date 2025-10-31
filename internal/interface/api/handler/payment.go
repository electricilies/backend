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

func (h *paymentHandler) ListMethods(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
