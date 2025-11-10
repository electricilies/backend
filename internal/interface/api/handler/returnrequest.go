package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnRequest interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type returnRequestHandler struct{}

func NewReturn() ReturnRequest { return &returnRequestHandler{} }

// GetReturn godoc
//
//	@Summary		Get return request by ID
//	@Description	Get return request details by ID
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Param			return_request_id	path		int	true	"Return Request ID"
//	@Success		200					{object}	response.ReturnRequest
//	@Failure		404					{object}	response.NotFoundError
//	@Failure		500					{object}	response.InternalServerError
//	@Router			/return-requests/{return_request_id} [get]
func (h *returnRequestHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListReturns godoc
//
//	@Summary		List all return requests
//	@Description	Get all return requests
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.ReturnRequest
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/return-requests [get]
func (h *returnRequestHandler) List(ctx *gin.Context) {
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
//	@Success		201		{object}	response.ReturnRequest
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		409		{object}	response.ConflictError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/return-requests [post]
func (h *returnRequestHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateReturnStatus godoc
//
//	@Summary		Update return status
//	@Description	Update the status of a return request
//	@Tags			Return
//	@Accept			json
//	@Produce		json
//	@Param			return_request_id	path		int									true	"Return Request ID"
//	@Param			status				body		request.UpdateReturnRequestStatus	true	"Update return status request"
//	@Success		204					{string}	string								"no content"
//	@Failure		400					{object}	response.BadRequestError
//	@Failure		404					{object}	response.NotFoundError
//	@Failure		409					{object}	response.ConflictError
//	@Failure		500					{object}	response.InternalServerError
//	@Router			/return-requests/{return_request_id}/status [put]
func (h *returnRequestHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
