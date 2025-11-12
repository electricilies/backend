package handler

import (
	"net/http"
	"strconv"

	"backend/internal/application"
	"backend/internal/domain/attribute"
	"backend/internal/interface/api/request"
	"backend/internal/interface/api/response"

	"github.com/gin-gonic/gin"
)

type Attribute interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UpdateAttributeValues(ctx *gin.Context)
}

type attributeHandler struct {
	app application.Attribute
}

func NewAttribute(app application.Attribute) Attribute {
	return &attributeHandler{
		app: app,
	}
}

// GetAttribute godoc
//
//	@Summary		Get attribute by ID
//	@Description	Get attribute details by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string	true	"Attribute ID"
//	@Success		200				{object}	response.Attribute
//	@Failure		404				{object}	response.NotFoundError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/attributes/{attribute_id} [get]
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
//	@Param			offset		query		int	false	"Offset for pagination"
//	@Param			limit		query		int	false	"Limit for pagination"
//	@Param			product_id	query		int	false	"Product ID"
//
//	@Success		200			{object}	response.DataPagination{data=[]response.Attribute}
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/attributes [get]
func (h *attributeHandler) List(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset")) // TODO: check, now it not required
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	productID, err := strconv.Atoi(ctx.Query("product_id"))
	if err != nil {
		productID = 0
	}
	pagination, err := h.app.ListAttributes(ctx, &attribute.QueryParams{
		PaginationParams: request.PaginationToDomain(limit, offset),
		ProductID:        &productID,
	})
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.DataPaginationFromDomain(pagination.Attributes, pagination.Metadata))
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
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		409			{object}	response.ConflictError
//	@Failure		500			{object}	response.InternalServerError
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
//	@Param			attribute_id	path		string					true	"Attribute ID"
//	@Param			attribute		body		request.UpdateAttribute	true	"Update attribute request"
//	@Success		204				{string}	string					"no content"
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		404				{object}	response.NotFoundError
//	@Failure		409				{object}	response.ConflictError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/attributes/{attribute_id} [put]
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
//	@Param			attribute_id	path		string	true	"Attribute ID"
//	@Success		204				{string}	string	"no content"
//	@Failure		404				{object}	response.NotFoundError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/attributes/{attribute_id} [delete]
func (h *attributeHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateAttributeValues godoc
//
//	@Summary		Update attribute values
//	@Description	Update attribute values for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path	string							true	"Attribute ID"
//	@Param			values			body	[]request.UpdateAttributeValue	true	"Update attribute values request"
//	@Success		204
//	@Failure		400	{object}	response.BadRequestError
//	@Failure		404	{object}	response.NotFoundError
//	@Failure		409	{object}	response.ConflictError
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/attributes/{attribute_id}/values/bulk [put]
func (h *attributeHandler) UpdateAttributeValues(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
