package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cart interface {
	Get(ctx *gin.Context)
	AddItem(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	RemoveItem(ctx *gin.Context)
}

type cartHandler struct{}

func NewCart() Cart { return &cartHandler{} }

func (h *cartHandler) Get(ctx *gin.Context)        { ctx.Status(http.StatusNoContent) }
func (h *cartHandler) AddItem(ctx *gin.Context)    { ctx.Status(http.StatusNoContent) }
func (h *cartHandler) UpdateItem(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
func (h *cartHandler) RemoveItem(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
