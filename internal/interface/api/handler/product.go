package handler

import (
	"net/http"
	"strconv"

	"backend/internal/application"
	"backend/internal/interface/api/request"
	"backend/internal/interface/api/response"

	"github.com/gin-gonic/gin"
)

type Product interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreateProductOption(ctx *gin.Context)
	CreateProductVariant(ctx *gin.Context)
	UpdateProductVariant(ctx *gin.Context)
	UpdateProductOption(ctx *gin.Context)
	GetDeleteImageURL(ctx *gin.Context)
	GetUploadImageURL(ctx *gin.Context)
	CreateProductImages(ctx *gin.Context)
}

type productHandler struct {
	app application.Product
}

func NewProduct(app application.Product) Product {
	return &productHandler{
		app: app,
	}
}

// GetProduct godoc
//
//	@Summary		Get product by ID
//	@Description	Get product details by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			produdt_id	path		int	true	"Product ID"
//	@Success		200			{object}	response.Product
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/products/{product_id} [get]
func (h *productHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListProducts godoc
//
//	@Summary		List all products
//	@Description	Get all products
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"Offset for pagination"
//	@Param			limit			query		int		false	"Limit for pagination"		default(20)
//	@Param			deleted			query		string	false	"Filter by deleted status"	Enums(exclude, only, all)
//	@Param			sort_price		query		string	false	"Sort by price"				Enums(asc, desc)
//	@Param			sort_rating		query		string	false	"Sort by rating"			Enums(asc, desc)
//	@Param			category_ids	query		[]int	false	"Filter by category ID"		CollectionFormat(csv)
//	@Param			min_price		query		int		false	"Minimum price filter"
//	@Param			max_price		query		int		false	"Maximum price filter"
//	@Success		200				{object}	response.DataPagination{data=[]response.Product}
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/products [get]
func (h *productHandler) List(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))   // TODO: check all pagination because now it not required
	offset, _ := strconv.Atoi(ctx.Query("offset")) // TODO: check if the json is checked of need to check in here
	pagination, error := h.app.ListProducts(ctx, request.ProductQueryParamsToDomain(&request.ProductQueryParams{
		Limit:      limit,
		Offset:     offset,
		Search:     ctx.Query("search"),
		Deleted:    ctx.Query("deleted"),
		SortPrice:  ctx.Query("sort_price"),
		SortRating: ctx.Query("sort_rating"),
		MinPrice: func() int64 {
			v, _ := strconv.ParseInt(ctx.Query("min_price"), 10, 64)
			return v
		}(),
		MaxPrice: func() int64 {
			v, _ := strconv.ParseInt(ctx.Query("max_price"), 10, 64)
			return v
		}(),
		CategoryIDs: func() []int {
			ids := []int{}
			for _, idStr := range ctx.QueryArray("category_ids") {
				id, _ := strconv.Atoi(idStr)
				ids = append(ids, id)
			}
			return ids
		}(),
	}))
	if error != nil {
		response.ErrorFromDomain(ctx, error)
		return
	}
	ctx.JSON(
		http.StatusOK,
		response.DataPaginationFromDomain(pagination.Products, pagination.Metadata),
	)
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product	body		request.CreateProduct	true	"Product request"
//	@Success		201		{object}	response.Product
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		409		{object}	response.ConflictError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/products [post]
func (h *productHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		int						true	"Product ID"
//	@Param			product		body		request.UpdateProduct	true	"Update product request"
//	@Success		200			{object}	response.Product
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		409			{object}	response.ConflictError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/products/{product_id} [patch]
func (h *productHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		int		true	"Product ID"
//	@Success		204			{string}	string	"no content"
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/products/{product_id} [delete]
func (h *productHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateProductOption godoc
//
//	@Summary		Create a new product option
//	@Description	Create a new product option for a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int							true	"Product ID"
//	@Param			productOption	body		request.CreateProductOption	true	"Product option request"
//
//	@Success		201				{object}	response.ProductOption
//
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		409				{object}	response.ConflictError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/products/options [post]
func (h *productHandler) CreateProductOption(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// GetUploadImageURL godoc
//
//	@Summary		Get presigned URL for image upload
//	@Description	Get a presigned URL to upload product images
//	@Tags			Product
//	@Produce		json
//	@Success		200	{object}	response.ProductUploadURLImage
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/products/images/upload-url [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *productHandler) GetUploadImageURL(ctx *gin.Context) {
	res, err := h.app.GetUploadImageURL(ctx)
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		response.ProductUploadURLImageFromDomain(res),
	)
}

// GetDeleteImageURL godoc
//
//	@Summary		Get presigned URL for image deletion
//	@Description	Get a presigned URL to delete product images
//	@Tags			Product
//	@Produce		json
//
//	@Param			image_id	query		int	true	"Product Image ID"
//
//	@Success		204			{object}	response.ProductImageDeleteURL
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/products/images/delete-url [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *productHandler) GetDeleteImageURL(ctx *gin.Context) {
	q := ctx.Query("image_id")
	id, _ := strconv.Atoi(q)
	url, err := h.app.GetDeleteImageURL(ctx, id)
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.ProductImageDeleteURL{URL: url},
	)
}

// CreateProductVariant godoc
//
//	@Summary		Create a new product variant
//	@Description	Create a new variant for a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int								true	"Product ID"``
//	@Param			productVariant	body		request.CreateProductVariant	true	"Product variant request"
//	@Success		201				{object}	response.ProductVariant
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		409				{object}	response.ConflictError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/products/variants [post]
func (h *productHandler) CreateProductVariant(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateProductVariant godoc
//
//	@Summary		Update a product variant
//	@Description	Update a product variant by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			variant_id		path		int								true	"Product Variant ID"
//	@Param			productVariant	body		request.UpdateProductVariant	true	"Update product variant request"
//	@Success		200				{object}	response.ProductVariant
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		404				{object}	response.NotFoundError
//	@Failure		409				{object}	response.ConflictError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/products/variants/{variant_id} [patch]
func (h *productHandler) UpdateProductVariant(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateProductOption godoc
//
//	@Summary		Update a product option
//	@Description	Update a product option by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int							true	"Product ID"
//	@Param			option_id		path		int							true	"Product Option ID"
//	@Param			productOption	body		request.UpdateProductOption	true	"Update product option request"
//	@Success		200				{object}	response.ProductOption
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		404				{object}	response.NotFoundError
//	@Failure		409				{object}	response.ConflictError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/products/options/{option_id} [patch]
func (h *productHandler) UpdateProductOption(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateProductImages godoc
//
//	@Summary		Create product images
//	@Description	Create new images for products
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Params			product_id   path        int                     true    "Product ID"
//	@Param			productImages	body		[]request.CreateProductImage	true	"Product images request"
//	@Success		201				{array}		response.ProductImage
//	@Failure		400				{object}	response.BadRequestError
//	@Failure		409				{object}	response.ConflictError
//	@Failure		500				{object}	response.InternalServerError
//	@Router			/products/{product_id}/images/bulk [post]
func (h *productHandler) CreateProductImages(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
