package http

import (
	"net/http"

	_ "backend/internal/domain"
	_ "backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AttributeHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	UpdateAttributeValues(*gin.Context)
}

type GinAttributeHandler struct{}

var _ AttributeHandler = &GinAttributeHandler{}

func ProvideAttributeHandler() *GinAttributeHandler {
	return &GinAttributeHandler{}
}

// GetAttribute godoc
//
//	@Summary		Get attribute by ID
//	@Description	Get attribute details by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string	true	"Attribute ID"
//	@Success		200				{object} domain.Attribute
//	@Failure		404				{object}	service.NotFoundError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/attributes/{attribute_id} [get]
func (h *GinAttributeHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListAttributes godoc
//
//	@Summary		List all attributes
//	@Description	Get all attributes
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int		false	"Offset for pagination"
//	@Param			limit		query		int		false	"Limit for pagination"	default(20)
//	@Param			product_id	query		int		false	"Product ID"
//	@Param			search		query		string	false	"Search term"
//	@Param			deleted		query		string	false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200			{object} domain.DataPagination{data=[]domain.Attribute}
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/attributes [get]
func (h *GinAttributeHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateAttribute godoc
//
//	@Summary		Create a new attribute
//	@Description	Create a new attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute	body service.CreateAttributeParam	true	"Attribute request"
//	@Success		201			{object}	domain.Attribute
//	@Failure		400			{object}	service.BadRequestError
//	@Failure		409			{object}	service.ConflictError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/attributes [post]
func (h *GinAttributeHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateAttribute godoc
//
//	@Summary		Update an attribute
//	@Description	Update attribute by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string					true	"Attribute ID"
//	@Param			attribute		body service.UpdateAttributeValueParam	true	"Update attribute request"
//	@Success		200				{object} domain.Attribute
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		404				{object}	service.NotFoundError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/attributes/{attribute_id} [patch]
func (h *GinAttributeHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteAttribute godoc
//
//	@Summary		Delete an attribute
//	@Description	Delete attribute by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string	true	"Attribute ID"
//	@Success		204
//	@Failure		404				{object}	service.NotFoundError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/attributes/{attribute_id} [delete]
func (h *GinAttributeHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateAttributeValues godoc
//
//	@Summary		Update attribute values
//	@Description	Update attribute values for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string							true	"Attribute ID"
//	@Param			values			body		[]service.UpdateAttributeValueParam	true	"Update attribute values request"
//	@Success		200				{array} domain.Attribute
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		404				{object}	service.NotFoundError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/attributes/{attribute_id}/values/bulk [patch]
func (h *GinAttributeHandler) UpdateAttributeValues(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
