package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Review interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type reviewHandler struct{}

func NewReview() Review { return &reviewHandler{} }

func (h *reviewHandler) Get(ctx *gin.Context)    { ctx.Status(http.StatusNoContent) }
func (h *reviewHandler) List(ctx *gin.Context)   { ctx.Status(http.StatusNoContent) }
func (h *reviewHandler) Create(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
func (h *reviewHandler) Update(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
func (h *reviewHandler) Delete(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
