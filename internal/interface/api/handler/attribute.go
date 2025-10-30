package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Attribute interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	ListValuesByAttribute(ctx *gin.Context)
	CreateValue(ctx *gin.Context)
}

type attributeHandler struct{}

func NewAttribute() Attribute { return &attributeHandler{} }

func (h *attributeHandler) Get(ctx *gin.Context)                   { ctx.Status(http.StatusNoContent) }
func (h *attributeHandler) List(ctx *gin.Context)                  { ctx.Status(http.StatusNoContent) }
func (h *attributeHandler) Create(ctx *gin.Context)                { ctx.Status(http.StatusNoContent) }
func (h *attributeHandler) Update(ctx *gin.Context)                { ctx.Status(http.StatusNoContent) }
func (h *attributeHandler) Delete(ctx *gin.Context)                { ctx.Status(http.StatusNoContent) }
func (h *attributeHandler) ListValuesByAttribute(ctx *gin.Context) { ctx.Status(http.StatusNoContent) }
func (h *attributeHandler) CreateValue(ctx *gin.Context)           { ctx.Status(http.StatusNoContent) }
