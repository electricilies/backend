package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Return interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
}

type returnHandler struct{}

func NewReturn() Return { return &returnHandler{} }

// GetReturn godoc
//
//	@Summary		Get return request by ID
//	@Description	Get return request details by ID
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Return Request ID"
//	@Success		200	{object}	response.Return
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/returns/{id} [get]
func (h *returnHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListReturns godoc
//
//	@Summary		List all return requests
//	@Description	Get all return requests
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Return
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/returns [get]
func (h *returnHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateReturn godoc
//
//	@Summary		Create a return request
//	@Description	Create a new return request
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Param			return	body		request.CreateReturnRequest	true	"Return request"
//	@Success		201		{object}	response.Return
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		409		{object}	mapper.ConflictError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/returns [post]
func (h *returnHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateReturnStatus godoc
//
//	@Summary		Update return status
//	@Description	Update the status of a return request
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Return Request ID"
//	@Param			status	body		request.UpdateReturnStatus	true	"Update return status request"
//	@Success		204		{string}	string						"no content"
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		404		{object}	mapper.NotFoundError
//	@Failure		409		{object}	mapper.ConflictError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/returns/{id}/status [put]
func (h *returnHandler) UpdateStatus(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
