package handler

import (
	"backend/internal/application"
	"backend/internal/interface/api/mapper"
	"backend/internal/interface/api/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreateProductOption(ctx *gin.Context)
	CreateProductImage(ctx *gin.Context)
	CreateProductVariant(ctx *gin.Context)
	UpdateProductVariant(ctx *gin.Context)
	UpdateProductOption(ctx *gin.Context)
	GetUploadImageURL(ctx *gin.Context)
	GetDeleteImageURL(ctx *gin.Context)
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
//	@Success		200			{object}	response.ProductWithVariants
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		500			{object}	mapper.InternalServerError
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
//	@Success		200	{array}		response.Product
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products [get]
func (h *productHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product	body		request.CreateProduct	true	"Product request"
//	@Success		201		{object}	response.ProductWithVariants
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		409		{object}	mapper.ConflictError
//	@Failure		500		{object}	mapper.InternalServerError
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
//	@Success		204			{string}	string					"no content"
//	@Failure		400			{object}	mapper.BadRequestError
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		409			{object}	mapper.ConflictError
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/products/{product_id} [put]
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
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		500			{object}	mapper.InternalServerError
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
//	@Param			productOption	body		request.CreateProductOption	true	"Product option request"
//
//	@Success		201				{object}	response.ProductOption
//
//	@Failure		400				{object}	mapper.BadRequestError
//	@Failure		409				{object}	mapper.ConflictError
//	@Failure		500				{object}	mapper.InternalServerError
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
//	@Success		200	{object}	response.ProductImageUploadURL
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products/images/upload-url [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2PasswordAdmin
//	@Security		OAuth2PasswordStaff
//	@Security		OAuth2PasswordCustomer
func (h *productHandler) GetUploadImageURL(ctx *gin.Context) {
	url, err := h.app.GetUploadImageURL(ctx)
	if err != nil {
		mapper.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.ProductImageUploadURL{URL: url},
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
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/products/images/delete-url [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2PasswordAdmin
//	@Security		OAuth2PasswordStaff
//	@Security		OAuth2PasswordCustomer
func (h *productHandler) GetDeleteImageURL(ctx *gin.Context) {
	q := ctx.Query("image_id")
	id, _ := strconv.Atoi(q)
	url, err := h.app.GetDeleteImageURL(ctx, id)
	if err != nil {
		mapper.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.ProductImageDeleteURL{URL: url},
	)
}

// CreateProductImage godoc
//
//	@Summary		Create a new product image
//	@Description	Create a new image for a product variant
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			productImage	body		request.CreateProductImage	true	"Product image request"
//	@Success		201				{object}	response.ProductImage
//	@Failure		400				{object}	mapper.BadRequestError
//	@Failure		409				{object}	mapper.ConflictError
//	@Failure		500				{object}	mapper.InternalServerError
//	@Router			/products/images [post]
func (h *productHandler) CreateProductImage(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateProductVariant godoc
//
//	@Summary		Create a new product variant
//	@Description	Create a new variant for a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			productVariant	body		request.CreateProductVariant	true	"Product variant request"
//	@Success		201				{object}	response.ProductVariantWithImages
//	@Failure		400				{object}	mapper.BadRequestError
//	@Failure		409				{object}	mapper.ConflictError
//	@Failure		500				{object}	mapper.InternalServerError
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
//	@Success		204				{string}	string							"no content"
//	@Failure		400				{object}	mapper.BadRequestError
//	@Failure		404				{object}	mapper.NotFoundError
//	@Failure		409				{object}	mapper.ConflictError
//	@Failure		500				{object}	mapper.InternalServerError
//	@Router			/products/variants/{variant_id} [put]
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
//	@Param			option_id		path		int							true	"Product Option ID"
//	@Param			productOption	body		request.UpdateProductOption	true	"Update product option request"
//	@Success		204				{string}	string						"no content"
//	@Failure		400				{object}	mapper.BadRequestError
//	@Failure		404				{object}	mapper.NotFoundError
//	@Failure		409				{object}	mapper.ConflictError
//	@Failure		500				{object}	mapper.InternalServerError
//	@Router			/products/options/{option_id} [put]
func (h *productHandler) UpdateProductOption(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
