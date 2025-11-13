package handler

import (
	"net/http"
	"strconv"

	"backend/internal/application"
	"backend/internal/domain/category"
	"backend/internal/interface/api/request"
	"backend/internal/interface/api/response"

	"github.com/gin-gonic/gin"
)

type Category interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type categoryHandler struct {
	app application.Category
}

func NewCategory(app application.Category) Category {
	return &categoryHandler{
		app: app,
	}
}

// ListCategories godoc
//
//	@Summary		List all categories
//	@Description	Get all categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int	false	"Offset for pagination"
//	@Param			limit	query		int	false	"Limit for pagination"
//	@Success		200		{object}	response.DataPagination{data=[]response.Category}
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/categories [get]
func (h *categoryHandler) List(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset")) // TODO: now it not required
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	pagination, err := h.app.ListCategories(ctx, &category.QueryParams{
		PaginationParams: request.PaginationToDomain(limit, offset),
	})
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.DataPaginationFromDomain(pagination.Categories, pagination.Metadata))
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	Create a new category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		request.CreateCategory	true	"Category request"
//	@Success		201			{object}	response.Category
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		409			{object}	response.ConflictError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/categories [post]
func (h *categoryHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	Update category by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category_id	path		int						true	"Category ID"
//	@Param			category	body		request.UpdateCategory	true	"Update category request"
//	@Success		204			{string}	string					"no content"
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		409			{object}	response.ConflictError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/categories/{category_id} [put]
func (h *categoryHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
