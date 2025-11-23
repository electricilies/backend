package http

import (
	"net/http"

	_ "backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type ProductHandlerImpl struct{}

var _ ProductHandler = &ProductHandlerImpl{}

func ProvideProductHandler() *ProductHandlerImpl {
	return &ProductHandlerImpl{}
}

// GetProduct godoc
//
//	@Summary		Get product by ID
//	@Description	Get product details by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		int	true	"Product ID"	format(uuid)
//	@Success		200			{object}	domain.Product
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id} [get]
func (h *ProductHandlerImpl) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListProducts godoc
//
//	@Summary		List all products
//	@Description	Get all products, used for search and suggestions also
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			page			query		int		false	"Page for pagination"		default(1)
//	@Param			limit			query		int		false	"Limit for pagination"		default(20)
//	@Param			deleted			query		string	false	"Filter by deleted status"	Enums(exclude, only, all)
//	@Param			sort_price		query		string	false	"Sort by price"				Enums(asc, desc)
//	@Param			sort_rating		query		string	false	"Sort by rating"			Enums(asc, desc)
//	@Param			category_ids	query		[]int	false	"Filter by category ID"		CollectionFormat(csv)	format(uuid)
//	@Param			min_price		query		int		false	"Minimum price filter"
//	@Param			max_price		query		int		false	"Maximum price filter"
//	@Success		200				{object}	application.Pagination[domain.Product]
//	@Failure		500				{object}	Error
//	@Router			/products [get]
func (h *ProductHandlerImpl) List(ctx *gin.Context) {
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product, including allllll
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product	body		application.CreateProductData	true	"Product request"
//	@Success		201		{object}	domain.Product
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/products [post]
func (h *ProductHandlerImpl) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		int								true	"Product ID"	format(uuid)
//	@Param			product		body		application.UpdateProductData	true	"Update product request"
//	@Success		200			{object}	domain.Product
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id} [patch]
func (h *ProductHandlerImpl) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path	int	true	"Product ID"	format(uuid)
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/products/{product_id} [delete]
func (h *ProductHandlerImpl) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// AddImages godoc
//
//	@Summary		Add product images
//	@Description	Create new images for an existing product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int										true	"Product ID"
//	@Param			productImages	body		[]application.CreateProductImageData	true	"Product images request"
//	@Success		201				{array}		domain.ProductImage
//	@Failure		400				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/images [post]
func (h *ProductHandlerImpl) AddImages(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteImages godoc
//
//	@Summary		Delete product images
//	@Description	Delete images for an existing product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path	int		true	"Product ID"	format(uuid)
//	@Param			ids			query	[]int	true	"Product Image IDs"
//	@Success		204
//	@Failure		400	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/products/{product_id}/images [delete]
func (h *ProductHandlerImpl) DeleteImages(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// AddVariants godoc
//
//	@Summary		Add a new product variant
//	@Description	Add a new variant for a existing product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int										true	"Product ID"
//	@Param			productVariant	body		[]application.CreateProductVariantData	true	"Product variant request"
//	@Success		201				{object}	domain.ProductVariant
//	@Failure		400				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/variants [post]
func (h *ProductHandlerImpl) AddVariants(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateVariant godoc
//
//	@Summary		Update a product variant
//	@Description	Update a product variant by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		int										true	"Product ID"
//	@Param			variant_id		path		int										true	"Product Variant ID"
//	@Param			productVariant	body		application.UpdateProductVariantData	true	"Update product variant request"
//	@Success		200				{object}	[]domain.ProductVariant
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/variants/{variant_id} [patch]
func (h *ProductHandlerImpl) UpdateVariant(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateOptions godoc
//
//	@Summary		Update options
//	@Description	Update a product options
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		int										true	"Product ID"	format(uuid)
//	@Param			option		body		[]application.UpdateProductOptionsData	true	"Update product option request"
//	@Success		200			{object}	domain.Option
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id}/options [put]
func (h *ProductHandlerImpl) UpdateOptions(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// GetUploadImageURL godoc
//
//	@Summary		Get presigned URL for image upload
//	@Description	Get a presigned URL to upload product images
//	@Tags			Product
//	@Produce		json
//	@Success		200	{object}	application.UploadImageURL
//	@Failure		500	{object}	Error
//	@Router			/products/images/upload-url [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) GetUploadImageURL(ctx *gin.Context) {
}

// GetDeleteImageURL godoc
//
//	@Summary		Get presigned URL for image deletion
//	@Description	Get a presigned URL to delete product images
//	@Tags			Product
//	@Produce		json
//	@Param			image_id	path		int	true	"Product Image ID"	format(uuid)
//	@Success		204			{object}	application.DeleteImageURL
//	@Failure		500			{object}	Error
//	@Router			/products/images/delete-url/{image_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) GetDeleteImageURL(ctx *gin.Context) {
}
