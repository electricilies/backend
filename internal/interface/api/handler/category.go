package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type categoryHandler struct{}

func NewCategory() Category { return &categoryHandler{} }

// ListCategories godoc
//
//	@Summary		List all categories
//	@Description	Get all categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.Category
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/categories [get]
func (h *categoryHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
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
//	@Failure		400			{object}	mapper.BadRequestError
//	@Failure		409			{object}	mapper.ConflictError
//	@Failure		500			{object}	mapper.InternalServerError
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
//	@Failure		400			{object}	mapper.BadRequestError
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		409			{object}	mapper.ConflictError
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/categories/{category_id} [put]
func (h *categoryHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteCategory godoc
//
//	@Summary		Delete a category
//	@Description	Delete category by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category_id	path		int		true	"Category ID"
//	@Success		204			{string}	string	"no content"
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/categories/{category_id} [delete]
func (h *categoryHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
