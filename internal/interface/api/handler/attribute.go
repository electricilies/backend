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
}

type attributeHandler struct{}

func NewAttribute() Attribute { return &attributeHandler{} }

// GetAttribute godoc
//
//	@Summary		Get attribute by ID
//	@Description	Get attribute details by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Attribute ID"
//	@Success		200	{object}	response.Attribute
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/attributes/{id} [get]
func (h *attributeHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListAttributes godoc
//
//	@Summary		List all attributes
//	@Description	Get all attributes
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Attribute
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/attributes [get]
func (h *attributeHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateAttribute godoc
//
//	@Summary		Create a new attribute
//	@Description	Create a new attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute	body		request.CreateAttribute	true	"Attribute request"
//	@Success		201			{object}	response.Attribute
//	@Failure		400			{object}	mapper.BadRequestError
//	@Failure		409			{object}	mapper.ConflictError
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/attributes [post]
func (h *attributeHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateAttribute godoc
//
//	@Summary		Update an attribute
//	@Description	Update attribute by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string					true	"Attribute ID"
//	@Param			attribute	body		request.UpdateAttribute	true	"Update attribute request"
//	@Success		204			{string}	string					"no content"
//	@Failure		400			{object}	mapper.BadRequestError
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		409			{object}	mapper.ConflictError
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/attributes/{id} [put]
func (h *attributeHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteAttribute godoc
//
//	@Summary		Delete an attribute
//	@Description	Delete attribute by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Attribute ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/attributes/{id} [delete]
func (h *attributeHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
