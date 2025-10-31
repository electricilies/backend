package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ListItemByOrder(ctx *gin.Context)
}

type orderHandler struct{}

func NewOrder() Order { return &orderHandler{} }

func (h *orderHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) UpdateStatus(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) ListItemByOrder(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
