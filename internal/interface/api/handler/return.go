package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Return interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
}

type returnHandler struct{}

func NewReturn() Return { return &returnHandler{} }

func (h *returnHandler) Get(ctx *gin.Context)          { ctx.Status(http.StatusNoContent) }
func (h *returnHandler) List(ctx *gin.Context)         { ctx.Status(http.StatusNoContent) }
func (h *returnHandler) Create(ctx *gin.Context)       { ctx.Status(http.StatusNoContent) }
func (h *returnHandler) UpdateStatus(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
