package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type categoryHandler struct{}

func NewCategory() Category { return &categoryHandler{} }

func (h *categoryHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *categoryHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *categoryHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *categoryHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *categoryHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
