package http

import (
	"net/http"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttributeHandlerImpl struct {
	attributeApp                AttributeApplication
	ErrRequiredAttributeID      string
	ErrRequiredAttributeValueID string
	ErrInvalidAttributeID       string
	ErrInvalidAttributeValueID  string
	ErrInvalidProductID         string
}

var _ AttributeHandler = &AttributeHandlerImpl{}

func ProvideAttributeHandler(attributeApp AttributeApplication) *AttributeHandlerImpl {
	return &AttributeHandlerImpl{
		attributeApp:                attributeApp,
		ErrRequiredAttributeID:      "attribute_id is required",
		ErrRequiredAttributeValueID: "attribute_value_id is required",
		ErrInvalidAttributeValueID:  "invalid value_id",
		ErrInvalidAttributeID:       "invalid attribute_id",
	}
}

// GetAttribute godoc
//
//	@Summary		Get attribute by ID
//	@Description	Get attribute details by ID
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string	true	"Attribute ID"	format(uuid)
//	@Success		200				{object}	domain.Attribute
//	@Failure		404				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) Get(ctx *gin.Context) {
	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}
	attribute, err := h.attributeApp.Get(ctx, GetAttributeRequestDto{
		AttributeID: *attributeID,
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
//	@Param			page			query		int			false	"Page for pagination"	default(1)
//	@Param			limit			query		int			false	"Limit for pagination"	default(20)
//	@Param			attribute_ids	query		[]string	false	"Attribute IDs"			collectionFormat(csv)	format(uuid)
//	@Param			product_ids		query		[]string	false	"Product IDs"			collectionFormat(csv)
//	@Param			search			query		string		false	"Search term"
//	@Param			deleted			query		string		false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200				{object}	Pagination[domain.Attribute]
//	@Failure		500				{object}	Error
//	@Router			/attributes [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) List(ctx *gin.Context) {
	paginateParam, err := createPaginationRequestDtoFromQuery(ctx)
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
	attributes, err := h.attributeApp.List(ctx, ListAttributesRequestDto{
		PaginationRequestDto: *paginateParam,
		AttributeIDs:         attributeIDs,
		Search:               search,
		Deleted:              deleted,
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
//	@Param			page				query		int			false	"Page for pagination"	default(1)
//	@Param			limit				query		int			false	"Limit for pagination"	default(20)
//	@Param			attribute_id		path		string		true	"Attribute ID"
//	@Param			attribute_value_id	query		[]string	false	"Product ID"	collectionFormat(csv)	format(uuid)
//	@Param			search				query		string		false	"Search term"
//	@Param			deleted				query		string		false	"Filter by deletion status"	Enums(exclude, only, all)
//	@Success		200					{object}	Pagination[domain.AttributeValue]
//	@Failure		500					{object}	Error
//	@Router			/attributes/{attribute_id}/values [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) ListValues(ctx *gin.Context) {
	paginateParam, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
	}

	var attributeValueIDs *[]uuid.UUID
	if attributeValueIDsQuery, ok := queryArrayToUUIDSlice(ctx, "attribute_value_ids"); ok {
		attributeValueIDs = attributeValueIDsQuery
	}

	var search *string
	if searchQuery, ok := ctx.GetQuery("search"); ok {
		search = &searchQuery
	}
	deleted := domain.DeletedExcludeParam
	if deletedQuery, ok := ctx.GetQuery("deleted"); ok {
		deleted = domain.DeletedParam(deletedQuery)
	}
	attributeValues, err := h.attributeApp.ListValues(ctx, ListAttributeValuesRequestDto{
		PaginationRequestDto: *paginateParam,
		AttributeID:          *attributeID,
		AttributeValueIDs:    attributeValueIDs,
		Deleted:              deleted,
		Search:               search,
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
//	@Param			attribute	body		CreateAttributeData	true	"Attribute request"
//	@Success		201			{object}	domain.Attribute
//	@Failure		400			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/attributes [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) Create(ctx *gin.Context) {
	var data CreateAttributeData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attribute, err := h.attributeApp.Create(ctx, CreateAttributeRequestDto{
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
//	@Param			attribute_id	path		string						true	"Attribute ID"	format(uuid)
//	@Param			value			body		CreateAttributeValueData	true	"Attribute value request"
//	@Success		201				{object}	domain.AttributeValue
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id}/values [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) CreateValue(ctx *gin.Context) {
	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	var data CreateAttributeValueData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attributeValue, err := h.attributeApp.CreateValue(ctx, CreateAttributeValueRequestDto{
		AttributeID: *attributeID,
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
//	@Param			attribute_id	path		string						true	"Attribute ID"	format(uuid)
//	@Param			attribute		body		UpdateAttributeValueData	true	"Update attribute request"
//	@Success		200				{object}	domain.Attribute
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) Update(ctx *gin.Context) {
	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	var data UpdateAttributeData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attribute, err := h.attributeApp.Update(ctx, UpdateAttributeRequestDto{
		AttributeID: *attributeID,
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
//	@Param			attribute_id	path	string	true	"Attribute ID"	format(uuid)
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/attributes/{attribute_id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) Delete(ctx *gin.Context) {
	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	err := h.attributeApp.Delete(ctx, DeleteAttributeRequestDto{
		AttributeID: *attributeID,
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
//	@Param			attribute_id	path	string	true	"Attribute ID"	format(uuid)
//	@Param			value_id		path	string	true	"Attribute Value ID"
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/attributes/{attribute_id}/values/{value_id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) DeleteValue(ctx *gin.Context) {
	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	valueID, ok := pathToUUID(ctx, "value_id")
	if *valueID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeValueID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeValueID))
		return
	}

	err := h.attributeApp.DeleteValue(ctx, DeleteAttributeValueRequestDto{
		AttributeID:      *attributeID,
		AttributeValueID: *valueID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// UpdateValues godoc
//
//	@Summary		Update a attribute value
//	@Description	Update a attribute value for a given attribute
//	@Tags			Attribute
//	@Accept			json
//	@Produce		json
//	@Param			attribute_id	path		string						true	"Attribute ID"			format(uuid)
//	@Param			value_id		path		string						true	"Attribute Value ID"	format(uuid)
//	@Param			value			body		UpdateAttributeValueData	true	"Update attribute values request"
//	@Success		200				{array}		domain.Attribute
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/attributes/{attribute_id}/values/{value_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *AttributeHandlerImpl) UpdateValue(ctx *gin.Context) {
	attributeID, ok := pathToUUID(ctx, "attribute_id")
	if *attributeID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeID))
		return
	}

	valueID, ok := pathToUUID(ctx, "value_id")
	if *valueID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredAttributeValueID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidAttributeValueID))
		return
	}
	var data UpdateAttributeValueData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	attributeValue, err := h.attributeApp.UpdateValue(ctx, UpdateAttributeValueRequestDto{
		AttributeID:      *attributeID,
		AttributeValueID: *valueID,
		Data:             data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, attributeValue)
}
