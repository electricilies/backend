package http

import (
	"net/http"

	_ "backend/internal/domain"
	_ "backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}

type GinCategoryHandler struct{}

var _ CategoryHandler = &GinCategoryHandler{}

func ProvideCategoryHandler() *GinCategoryHandler {
	return &GinCategoryHandler{}
}

// ListCategories godoc
//
//	@Summary		List all categories
//	@Description	Get all categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Page for pagination"
//	@Param			limit	query		int	false	"Limit for pagination"	default(20)
//	@Success		200		{object} domain.DataPagination{data=[]domain.Category}
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/categories [get]
func (h *GinCategoryHandler) List(ctx *gin.Context) {
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	Create a new category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		service.CreateCategoryParam	true	"Category request"
//	@Success		201			{object} domain.Category
//	@Failure		400			{object}	service.BadRequestError
//	@Failure		409			{object}	service.ConflictError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/categories [post]
func (h *GinCategoryHandler) Create(ctx *gin.Context) {
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
//	@Param			category	body		service.UpdateCategoryParam	true	"Update category request"
//	@Success		200			{object} domain.Category
//	@Failure		400			{object}	service.BadRequestError
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		409			{object}	service.ConflictError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/categories/{category_id} [patch]
func (h *GinCategoryHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
