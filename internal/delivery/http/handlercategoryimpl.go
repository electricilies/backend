package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandlerImpl struct {
	categoryApp           CategoryApplication
	ErrRequiredCategoryID string
	ErrInvalidCategoryID  string
}

var _ CategoryHandler = &CategoryHandlerImpl{}

func ProvideCategoryHandler(categoryApp CategoryApplication) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{
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
//	@Param			search	query		string	false	"Search term"
//	@Param			page	query		int		false	"Page for pagination"	default(1)
//	@Param			limit	query		int		false	"Limit for pagination"	default(20)
//	@Success		200		{object}	PaginationResponseDto[CategoryResponseDto]
//	@Failure		500		{object}	Error
//	@Router			/categories [get]
func (h *CategoryHandlerImpl) List(ctx *gin.Context) {
	paginateParam, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	search, _ := ctx.GetQuery("search")

	categories, err := h.categoryApp.List(ctx, ListCategoryRequestDto{
		PaginationRequestDto: *paginateParam,
		Search:               search,
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
//	@Param			category_id	path		string	true	"Category ID"	format(uuid)
//	@Success		200			{object}	CategoryResponseDto
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories/{category_id} [get]
func (h *CategoryHandlerImpl) Get(ctx *gin.Context) {
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
	category, err := h.categoryApp.Get(ctx, GetCategoryRequestDto{
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
//	@Param			category	body		CreateCategoryData	true	"Category request"
//	@Success		201			{object}	CategoryResponseDto
//	@Failure		400			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Create(ctx *gin.Context) {
	var data CreateCategoryData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	category, err := h.categoryApp.Create(ctx, CreateCategoryRequestDto{
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
//	@Param			category_id	path		string				true	"Category ID"	format(uuid)
//	@Param			category	body		UpdateCategoryData	true	"Update category request"
//	@Success		200			{object}	CategoryResponseDto
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories/{category_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Update(ctx *gin.Context) {
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

	var data UpdateCategoryData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	category, err := h.categoryApp.Update(ctx, UpdateCategoryRequestDto{
		CategoryID: categoryID,
		Data:       data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, category)
}
