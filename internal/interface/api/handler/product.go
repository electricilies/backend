package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	ListVariantsByProduct(ctx *gin.Context)
	ListImagesByProduct(ctx *gin.Context)
	ListReviewByProduct(ctx *gin.Context)
	ListAtributesByProduct(ctx *gin.Context)
	AddAttributeValues(ctx *gin.Context)
	RemoveAttributeValue(ctx *gin.Context)
}

type productHandler struct{}

func NewProduct() Product { return &productHandler{} }

func (h *productHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) ListVariantsByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) ListImagesByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) ListReviewByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) ListAtributesByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) AddAttributeValues(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (h *productHandler) RemoveAttributeValue(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
