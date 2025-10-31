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

// GetRefund godoc
//
// @Summary      Get refund by ID
// @Description  Get refund details by ID
// @Tags         Refund
// @Accept       json
// @Produce      json
// @Param        id path int true "Refund ID"
// @Success      200 {object} response.Refund
// @Failure      404 {object} mapper.NotFoundError
// @Failure      500 {object} mapper.InternalServerError
// @Router       /refunds/{id} [get]
func (h *refundHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListRefunds godoc
//
// @Summary      List all refunds
// @Description  Get all refunds
// @Tags         Refund
// @Accept       json
// @Produce      json
// @Success      200 {array} response.Refund
// @Failure      500 {object} mapper.InternalServerError
// @Router       /refunds [get]
func (h *refundHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
