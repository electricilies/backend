package http

import (
	"net/http"

	"backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler interface {
	List(*gin.Context)
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}

type GinCategoryHandler struct {
	categoryApp           application.Category
	ErrRequiredCategoryID string
	ErrInvalidCategoryID  string
}

var _ CategoryHandler = &GinCategoryHandler{}

func ProvideCategoryHandler(categoryApp application.Category) *GinCategoryHandler {
	return &GinCategoryHandler{
		categoryApp:           categoryApp,
		ErrRequiredCategoryID: "category_id is required",
		ErrInvalidCategoryID:  "invalid category_id",
	}
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
//	@Success		200		{object}	application.Pagination[domain.Category]
//	@Failure		500		{object}	Error
//	@Router			/categories [get]
func (h *GinCategoryHandler) List(ctx *gin.Context) {
	paginateParam, err := createPaginationParamsFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	var search *string
	if searchQuery, ok := ctx.GetQuery("search"); ok {
		search = &searchQuery
	}

	categories, err := h.categoryApp.List(ctx, application.ListCategoryParam{
		PaginationParam: *paginateParam,
		Search:          search,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// GetCategory godoc
//
//	@Summary		Get category by ID
//	@Description	Get category details by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category_id	path		int	true	"Category ID"
//	@Success		200			{object}	domain.Category
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories/{category_id} [get]
func (h *GinCategoryHandler) Get(ctx *gin.Context) {
	categoryIDString := ctx.Param("category_id")
	if categoryIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCategoryID))
		return
	}
	categoryID, err := uuid.Parse(categoryIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCategoryID))
		return
	}
	category, err := h.categoryApp.Get(ctx, application.GetCategoryParam{
		CategoryID: categoryID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	Create a new category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		application.CreateCategoryData	true	"Category request"
//	@Success		201			{object}	domain.Category
//	@Failure		400			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories [post]
func (h *GinCategoryHandler) Create(ctx *gin.Context) {
	var data application.CreateCategoryData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	category, err := h.categoryApp.Create(ctx, application.CreateCategoryParam{
		Data: data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	Update category by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category_id	path		int								true	"Category ID"
//	@Param			category	body		application.UpdateCategoryData	true	"Update category request"
//	@Success		200			{object}	domain.Category
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories/{category_id} [patch]
func (h *GinCategoryHandler) Update(ctx *gin.Context) {
	categoryIDString := ctx.Param("category_id")
	if categoryIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCategoryID))
		return
	}
	categoryID, err := uuid.Parse(categoryIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCategoryID))
		return
	}

	var data application.UpdateCategoryData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	category, err := h.categoryApp.Update(ctx, application.UpdateCategoryParam{
		CategoryID: categoryID,
		Data:       data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, category)
}
