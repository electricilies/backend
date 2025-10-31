package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Refund interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
}

type refundHandler struct{}

func NewRefund() Refund { return &refundHandler{} }

func (h *refundHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *refundHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
