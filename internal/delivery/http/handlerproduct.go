package http

import (
	"net/http"

	_ "backend/internal/domain"
	_ "backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	CreateProductOption(*gin.Context)
	CreateProductVariant(*gin.Context)
	UpdateProductVariant(*gin.Context)
	UpdateProductOption(*gin.Context)
	GetDeleteImageURL(*gin.Context)
	GetUploadImageURL(*gin.Context)
	CreateProductImages(*gin.Context)
}

type GinProductHandler struct{}

var _ ProductHandler = &GinProductHandler{}

func ProvideProductHandler() *GinProductHandler {
	return &GinProductHandler{}
}

// GetProduct godoc
//
//	@Summary		Get product by ID
//	@Description	Get product details by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			produdt_id	path		int	true	"Product ID"
//	@Success		200			{object} domain.Product
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/products/{product_id} [get]
func (h *GinProductHandler) Get(ctx *gin.Context) {
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
//	@Success		200				{object} domain.DataPagination{data=[]domain.Product}
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/products [get]
func (h *GinProductHandler) List(ctx *gin.Context) {
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product	body service.CreateProductParam	true	"Product request"
//	@Success		201		{object} domain.Product
//	@Failure		400		{object}	service.BadRequestError
//	@Failure		409		{object}	service.ConflictError
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/products [post]
func (h *GinProductHandler) Create(ctx *gin.Context) {
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
//	@Param			product		body		service.UpdateProductParam	true	"Update product request"
//	@Success		200			{object} domain.Product
//	@Failure		400			{object}	service.BadRequestError
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		409			{object}	service.ConflictError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/products/{product_id} [patch]
func (h *GinProductHandler) Update(ctx *gin.Context) {
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
//	@Success		204
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/products/{product_id} [delete]
func (h *GinProductHandler) Delete(ctx *gin.Context) {
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
//	@Param			productOption	body service.CreateProductOptionParam	true	"Product option request"
//
//	@Success		201				{object} domain.ProductOption
//
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/products/options [post]
func (h *GinProductHandler) CreateProductOption(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// GetUploadImageURL godoc
//
//	@Summary		Get presigned URL for image upload
//	@Description	Get a presigned URL to upload product images
//	@Tags			Product
//	@Produce		json
//	@Success		200	{object} domain.ProductUploadURLImage
//	@Failure		500	{object}	service.InternalServerError
//	@Router			/products/images/upload-url [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *GinProductHandler) GetUploadImageURL(ctx *gin.Context) {
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
//	@Success		204			{object} domain.ProductImageDeleteURL
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/products/images/delete-url [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *GinProductHandler) GetDeleteImageURL(ctx *gin.Context) {
}

// CreateProductVariant godoc
//
//	@Summary		Create a new product variant
//	@Description	Create a new variant for a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int								true	"Product ID"``
//	@Param			productVariant	body service.CreateProductVariantParam	true	"Product variant request"
//	@Success		201				{object} domain.ProductVariant
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/products/variants [post]
func (h *GinProductHandler) CreateProductVariant(ctx *gin.Context) {
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
//	@Param			productVariant	body service.UpdateProductVariantParam	true	"Update product variant request"
//	@Success		200				{object} domain.ProductVariant
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		404				{object}	service.NotFoundError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/products/variants/{variant_id} [patch]
func (h *GinProductHandler) UpdateProductVariant(ctx *gin.Context) {
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
//	@Param			productOption	body service.UpdateProductOptionParam true	"Update product option request"
//	@Success		200				{object} domain.ProductOption
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		404				{object}	service.NotFoundError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/products/options/{option_id} [patch]
func (h *GinProductHandler) UpdateProductOption(ctx *gin.Context) {
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
//	@Param			productImages	body		[]service.CreateProductImageParam	true	"Product images request"
//	@Success		201				{array} domain.ProductImage
//	@Failure		400				{object}	service.BadRequestError
//	@Failure		409				{object}	service.ConflictError
//	@Failure		500				{object}	service.InternalServerError
//	@Router			/products/{product_id}/images/bulk [post]
func (h *GinProductHandler) CreateProductImages(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
