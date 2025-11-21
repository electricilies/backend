package http

import (
	"net/http"

	_ "backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type AttributeHandler interface {
	Get(*gin.Context)
	ListValues(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	CreateValue(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	DeleteValue(*gin.Context)
	UpdateValue(*gin.Context)
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
//	@Success		200				{object}	domain.Attribute
//	@Failure		404				{object}	Error
//	@Failure		500				{object}	Error
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
//	@Param			page		query		int		false	"Page for pagination"
//	@Param			limit		query		int		false	"Limit for pagination"	default(20)
//	@Param			product_id	query		int		false	"Product ID"
//	@Param			search		query		string	false	"Search term"
//	@Param			deleted		query		string	false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200			{object}	application.Pagination[domain.Attribute]
//	@Failure		500			{object}	Error
//	@Router			/attributes [get]
func (h *GinAttributeHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListAttributeValues godoc
//
//	@Summary		List all attribute values
//	@Description	Get all attribute values
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int		false	"Page for pagination"
//	@Param			limit		query		int		false	"Limit for pagination"	default(20)
//	@Param			attribute_id	path string false	"Attribute ID"
//	@Param 			attribute_value_id	query		string false	"Product ID"
//	@Param			search		query		string	false	"Search term"
//	@Param			deleted		query		string	false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200			{object}	application.Pagination[domain.AttributeValue]
//	@Failure		500			{object}	Error
//	@Router			/attributes/values [get]
func (h *GinAttributeHandler) ListValues(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateAttribute godoc
//
//	@Summary		Create a new attribute
//	@Description	Create a new attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute	body		application.CreateAttributeData	true	"Attribute request"
//	@Success		201			{object}	domain.Attribute
//	@Failure		400			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/attributes [post]
func (h *GinAttributeHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateAttributeValue godoc
//
//	@Summary		Create a new attribute value
//	@Description	Create a new attribute value for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string							true	"Attribute ID"
//	@Param			value			body		application.CreateAttributeValueData	true	"Attribute value request"
//	@Success		201				{object}	domain.AttributeValue
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id}/values [post]
func (h *GinAttributeHandler) CreateValue(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateAttribute godoc
//
//	@Summary		Update an attribute
//	@Description	Update attribute by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string							true	"Attribute ID"
//	@Param			attribute		body		application.UpdateAttributeValueData	true	"Update attribute request"
//	@Success		200				{object}	domain.Attribute
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
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
//	@Param			attribute_id	path	string	true	"Attribute ID"
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/attributes/{attribute_id} [delete]
func (h *GinAttributeHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteAttributeValue godoc
//
//	@Summary		Delete an attribute value
//	@Description	Delete attribute value by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			value_id	path	string	true	"Attribute Value ID"
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/attributes/values/{value_id} [delete]
func (h *GinAttributeHandler) DeleteValue(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateValues godoc
//
//	@Summary		Update attribute values
//	@Description	Update attribute values for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string								true	"Attribute ID"
//	@Param			values			body		[]application.UpdateAttributeValueData	true	"Update attribute values request"
//	@Success		200				{array}		domain.Attribute
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id}/values [patch]
func (h *GinAttributeHandler) UpdateValue(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
