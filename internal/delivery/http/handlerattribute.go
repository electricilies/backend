package http

import (
	"net/http"

	"backend/internal/application"
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

type GinAttributeHandler struct {
	attributeApp           application.Attribute
	ErrRequiredAttributeID string
	ErrInvalidAttributeID  string
	ErrInvalidProductID    string
}

var _ AttributeHandler = &GinAttributeHandler{}

func ProvideAttributeHandler(attributeApp application.Attribute) *GinAttributeHandler {
	return &GinAttributeHandler{
		attributeApp:           attributeApp,
		ErrRequiredAttributeID: "attribute_id is required",
		ErrInvalidAttributeID:  "invalid attribute_id",
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
//	@Success		200				{object}	domain.Attribute
//	@Failure		404				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id} [get]
func (h *GinAttributeHandler) Get(ctx *gin.Context) {
	attributeIDString := ctx.Param("attribute_id")
	if attributeIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	attributeID, err := uuid.Parse(attributeIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}
	attribute, err := h.attributeApp.Get(ctx, application.GetAttributeParam{
		AttributeID: attributeID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, attribute)
}

// ListAttributes godoc
//
//	@Summary		List all attributes
//	@Description	Get all attributes
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			page			query		int			false	"Page for pagination"
//	@Param			limit			query		int			false	"Limit for pagination"	default(20)
//	@Param			attribute_ids	query		[]string	false	"Attribute IDs"			collectionFormat(csv)
//	@Param			product_ids		query		[]string	false	"Product IDs"			collectionFormat(csv)
//	@Param			search			query		string		false	"Search term"
//	@Param			deleted			query		string		false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200				{object}	application.Pagination[domain.Attribute]
//	@Failure		500				{object}	Error
//	@Router			/attributes [get]
func (h *GinAttributeHandler) List(ctx *gin.Context) {
	paginateParam, err := createPaginationParamsFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	var attributeIDs *[]uuid.UUID
	if attributeIDsQuery, ok := queryArrayToUUIDSlice(ctx, "attribute_ids"); ok {
		attributeIDs = attributeIDsQuery
	}
	var search *string
	if searchQuery, ok := ctx.GetQuery("search"); ok {
		search = &searchQuery
	}
	deleted := domain.DeletedExcludeParam
	if deletedQuery, ok := ctx.GetQuery("deleted"); ok {
		deleted = domain.DeletedParam(deletedQuery)
	}
	attributes, err := h.attributeApp.List(ctx, application.ListAttributesParam{
		PaginationParam: *paginateParam,
		AttributeIDs:    attributeIDs,
		Search:          search,
		Deleted:         domain.DeletedParam(deleted),
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, attributes)
}

// ListAttributeValues godoc
//
//	@Summary		List all attribute values
//	@Description	Get all attribute values
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			page				query		int		false	"Page for pagination"
//	@Param			limit				query		int		false	"Limit for pagination"	default(20)
//	@Param			attribute_id		path		string	false	"Attribute ID"
//	@Param			attribute_value_id	query		string	false	"Product ID"
//	@Param			search				query		string	false	"Search term"
//	@Param			deleted				query		string	false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200					{object}	application.Pagination[domain.AttributeValue]
//	@Failure		500					{object}	Error
//	@Router			/attributes/values [get]
func (h *GinAttributeHandler) ListValues(ctx *gin.Context) {
	paginateParam, err := createPaginationParamsFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	var attributeID *uuid.UUID
	if attributeIDQuery, ok := ctx.GetQuery("attribute_id"); ok {
		parsedID, err := uuid.Parse(attributeIDQuery)
		if err == nil {
			attributeID = &parsedID
		}
	}

	var attributeValueIDs *[]uuid.UUID
	if attributeValueIDsQuery, ok := queryArrayToUUIDSlice(ctx, "attribute_value_ids"); ok {
		attributeValueIDs = attributeValueIDsQuery
	}

	var search *string
	if searchQuery, ok := ctx.GetQuery("search"); ok {
		search = &searchQuery
	}

	attributeValues, err := h.attributeApp.ListValues(ctx, application.ListAttributeValuesParam{
		PaginationParam:   *paginateParam,
		AttributeID:       attributeID,
		AttributeValueIDs: attributeValueIDs,
		Search:            search,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, attributeValues)
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
	var data application.CreateAttributeData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attribute, err := h.attributeApp.Create(ctx, application.CreateAttributeParam{
		Data: data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, attribute)
}

// CreateAttributeValue godoc
//
//	@Summary		Create a new attribute value
//	@Description	Create a new attribute value for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string									true	"Attribute ID"
//	@Param			value			body		application.CreateAttributeValueData	true	"Attribute value request"
//	@Success		201				{object}	domain.AttributeValue
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id}/values [post]
func (h *GinAttributeHandler) CreateValue(ctx *gin.Context) {
	attributeIDString := ctx.Param("attribute_id")
	if attributeIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	attributeID, err := uuid.Parse(attributeIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	var data application.CreateAttributeValueData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attributeValue, err := h.attributeApp.CreateValue(ctx, application.CreateAttributeValueParam{
		AttributeID: attributeID,
		Data:        data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, attributeValue)
}

// UpdateAttribute godoc
//
//	@Summary		Update an attribute
//	@Description	Update attribute by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string									true	"Attribute ID"
//	@Param			attribute		body		application.UpdateAttributeValueData	true	"Update attribute request"
//	@Success		200				{object}	domain.Attribute
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id} [patch]
func (h *GinAttributeHandler) Update(ctx *gin.Context) {
	attributeIDString := ctx.Param("attribute_id")
	if attributeIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	attributeID, err := uuid.Parse(attributeIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	var data application.UpdateAttributeData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attribute, err := h.attributeApp.Update(ctx, application.UpdateAttributeParam{
		AttributeID: attributeID,
		Data:        data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, attribute)
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
	attributeIDString := ctx.Param("attribute_id")
	if attributeIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	attributeID, err := uuid.Parse(attributeIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	err = h.attributeApp.Delete(ctx, application.DeleteAttributeParam{
		AttributeID: attributeID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
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
	valueIDString := ctx.Param("value_id")
	if valueIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError("value_id is required"))
		return
	}
	valueID, err := uuid.Parse(valueIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError("invalid value_id"))
		return
	}

	attributeIDString := ctx.Param("attribute_id")
	if attributeIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	attributeID, err := uuid.Parse(attributeIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	err = h.attributeApp.DeleteValue(ctx, application.DeleteAttributeValueParam{
		AttributeID:      attributeID,
		AttributeValueID: valueID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// UpdateValues godoc
//
//	@Summary		Update attribute values
//	@Description	Update attribute values for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string									true	"Attribute ID"
//	@Param			values			body		[]application.UpdateAttributeValueData	true	"Update attribute values request"
//	@Success		200				{array}		domain.Attribute
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id}/values [patch]
func (h *GinAttributeHandler) UpdateValue(ctx *gin.Context) {
	attributeIDString := ctx.Param("attribute_id")
	if attributeIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	attributeID, err := uuid.Parse(attributeIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	valueIDString := ctx.Param("value_id")
	if valueIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError("value_id is required"))
		return
	}
	valueID, err := uuid.Parse(valueIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError("invalid value_id"))
		return
	}

	var data application.UpdateAttributeValueData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attributeValue, err := h.attributeApp.UpdateValue(ctx, application.UpdateAttributeValueParam{
		AttributeID:      attributeID,
		AttributeValueID: valueID,
		Data:             data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, attributeValue)
}
