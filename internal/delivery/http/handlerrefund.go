package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefundHandler interface {
	// Get(*gin.Context)
	// List(*gin.Context)
}

type GinRefundHandler struct{}

var _ RefundHandler = &GinRefundHandler{}

func ProvideRefund() *GinRefundHandler {
	return &GinRefundHandler{}
}

// GetRefund godoc
//
//	@Summary		Get refund by ID
//	@Description	Get refund details by ID
//	@Tags			Refund
//	@Accept			json
//	@Produce		json
//	@Param			refund_id	path		int	true	"Refund ID"
//	@Success		200			{object}	response.Refund
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/refunds/{refund} [get]
// func (h *GinRefundHandler) Get(ctx *gin.Context) {
// 	ctx.Status(http.StatusNoContent)
// }

// ListRefunds godoc
//
//	@Summary		List all refunds
//	@Description	Get all refunds
//	@Tags			Refund
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Refund
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/refunds [get]
// func (h *GinRefundHandler) List(ctx *gin.Context) {
// 	ctx.Status(http.StatusNoContent)
// }
