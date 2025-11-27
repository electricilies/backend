package http

import (
	"fmt"
	"net/http"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type ProductHandlerImpl struct {
	productApp           ProductApplication
	ErrRequiredProductID string
	ErrInvalidProductID  string
}

var _ ProductHandler = &ProductHandlerImpl{}

func ProvideProductHandler(productApp ProductApplication) *ProductHandlerImpl {
	return &ProductHandlerImpl{
		productApp:           productApp,
		ErrRequiredProductID: "product_id is required",
		ErrInvalidProductID:  "invalid product_id",
	}
}

// GetProduct godoc
//
//	@Summary		Get product by ID
//	@Description	Get product details by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string	true	"Product ID"	format(uuid)
//	@Success		200			{object}	ProductResponseDto
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id} [get]
func (h *ProductHandlerImpl) Get(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}
	product, err := h.productApp.Get(ctx.Request.Context(), GetProductRequestDto{
		ProductID: productID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// ListProducts godoc
//
//	@Summary		List all products
//	@Description	Get all products, used for search and suggestions also
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			search			query		string		false	"Search term"
//	@Param			page			query		int			false	"Page for pagination"		default(1)
//	@Param			limit			query		int			false	"Limit for pagination"		default(20)
//	@Param			deleted			query		string		false	"Filter by deleted status"	Enums(exclude, only, all)
//	@Param			sort_price		query		string		false	"Sort by price"				Enums(asc, desc)
//	@Param			sort_rating		query		string		false	"Sort by rating"			Enums(asc, desc)
//	@Param			category_ids	query		[]string	false	"Filter by category ID"		CollectionFormat(csv)	format(uuid)
//	@Param			product_ids		query		[]string	false	"Filter by product ID"		CollectionFormat(csv)	format(uuid)
//	@Param			min_price		query		int			false	"Minimum price filter"
//	@Param			max_price		query		int			false	"Maximum price filter"
//	@Param			rating			query		number		false	"Filter by minimum rating"
//	@Success		200				{object}	PaginationResponseDto[ProductResponseDto]
//	@Failure		500				{object}	Error
//	@Router			/products [get]
func (h *ProductHandlerImpl) List(ctx *gin.Context) {
	paginateParam, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	productIDs, _ := queryArrayToUUIDSlice(ctx, "product_ids")

	categoryIDs, _ := queryArrayToUUIDSlice(ctx, "category_ids")

	search, _ := ctx.GetQuery("search")

	var minPrice int64
	if minPriceQuery, ok := ctx.GetQuery("min_price"); ok {
		var price int64
		if _, err := fmt.Sscanf(minPriceQuery, "%d", &price); err == nil {
			minPrice = price
		}
	}

	var maxPrice int64
	if maxPriceQuery, ok := ctx.GetQuery("max_price"); ok {
		var price int64
		if _, err := fmt.Sscanf(maxPriceQuery, "%d", &price); err == nil {
			maxPrice = price
		}
	}

	var rating float64
	if ratingQuery, ok := ctx.GetQuery("rating"); ok {
		var r float64
		if _, err := fmt.Sscanf(ratingQuery, "%f", &r); err == nil {
			rating = r
		}
	}

	sortPrice, _ := ctx.GetQuery("sort_price")

	sortRating, _ := ctx.GetQuery("sort_rating")

	deleted := domain.DeletedExcludeParam
	if deletedQuery, ok := ctx.GetQuery("deleted"); ok {
		deleted = domain.DeletedParam(deletedQuery)
	}

	products, err := h.productApp.List(ctx.Request.Context(), ListProductRequestDto{
		PaginationRequestDto: *paginateParam,
		ProductIDs:           productIDs,
		CategoryIDs:          categoryIDs,
		MinPrice:             minPrice,
		MaxPrice:             maxPrice,
		Rating:               rating,
		SortPrice:            sortPrice,
		SortRating:           sortRating,
		Search:               search,
		Deleted:              deleted,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product, including allllll
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product	body		CreateProductData	true	"Product request"
//	@Success		201		{object}	ProductResponseDto
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/products [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) Create(ctx *gin.Context) {
	var data CreateProductData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product, err := h.productApp.Create(ctx.Request.Context(), CreateProductRequestDto{
		Data: data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string				true	"Product ID"	format(uuid)
//	@Param			product		body		UpdateProductData	true	"Update product request"
//	@Success		200			{object}	ProductResponseDto
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) Update(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data UpdateProductData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product, err := h.productApp.Update(ctx.Request.Context(), UpdateProductRequestDto{
		ProductID: productID,
		Data:      data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path	string	true	"Product ID"	format(uuid)
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/products/{product_id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) Delete(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	err := h.productApp.Delete(ctx.Request.Context(), DeleteProductRequestDto{
		ProductID: productID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// AddImages godoc
//
//	@Summary		Add product images
//	@Description	Create new images for an existing product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		string					true	"Product ID"	format(uuid)
//	@Param			productImages	body		[]AddProductImageData	true	"Product images request"
//	@Success		201				{array}		ProductImageResponseDto
//	@Failure		400				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/images [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) AddImages(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data []AddProductImageData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	images, err := h.productApp.AddImages(ctx.Request.Context(), AddProductImagesRequestDto{
		ProductID: productID,
		Data:      data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, images)
}

// DeleteImages godoc
//
//	@Summary		Delete product images
//	@Description	Delete images for an existing product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path	string		true	"Product ID"		format(uuid)
//	@Param			ids			query	[]string	true	"Product Image IDs"	CollectionFormat(csv)	format(uuid)
//	@Success		204
//	@Failure		400	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/products/{product_id}/images [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) DeleteImages(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	imageIDs, ok := queryArrayToUUIDSlice(ctx, "ids")
	if !ok || len(imageIDs) == 0 {
		ctx.JSON(http.StatusBadRequest, NewError("image ids are required"))
		return
	}

	err := h.productApp.DeleteImages(ctx.Request.Context(), DeleteProductImagesRequestDto{
		ProductID: productID,
		ImageIDs:  imageIDs,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// AddVariants godoc
//
//	@Summary		Add a new product variant
//	@Description	Add a new variant for a existing product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		string						true	"Product ID"	format(uuid)
//	@Param			productVariant	body		[]AddProductVariantsData	true	"Product variant request"
//	@Success		201				{array}		ProductVariantResponseDto
//	@Failure		400				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/variants [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) AddVariants(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data []AddProductVariantsData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	variants, err := h.productApp.AddVariants(ctx.Request.Context(), AddProductVariantsRequestDto{
		ProductID: productID,
		Data:      data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, variants)
}

// UpdateVariant godoc
//
//	@Summary		Update a product variant
//	@Description	Update a product variant by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id		path		string						true	"Product ID"			format(uuid)
//	@Param			variant_id		path		string						true	"Product Variant ID"	format(uuid)
//	@Param			productVariant	body		UpdateProductVariantData	true	"Update product variant request"
//	@Success		200				{object}	ProductVariantResponseDto
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/variants/{variant_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) UpdateVariant(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	variantID, ok := pathToUUID(ctx, "variant_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError("invalid variant_id"))
		return
	}

	var data UpdateProductVariantData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	variant, err := h.productApp.UpdateVariant(ctx.Request.Context(), UpdateProductVariantRequestDto{
		ProductID:        productID,
		ProductVariantID: variantID,
		Data:             data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, variant)
}

// UpdateOptions godoc
//
//	@Summary		Update options
//	@Description	Update a product options
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string						true	"Product ID"	format(uuid)
//	@Param			option		body		[]UpdateProductOptionsData	true	"Update product option request"
//	@Success		200			{array}		ProductOptionResponseDto
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id}/options [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) UpdateOptions(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data []UpdateProductOptionsData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	options, err := h.productApp.UpdateOptions(ctx.Request.Context(), UpdateProductOptionsRequestDto{
		ProductID: productID,
		Data:      data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, options)
}

// GetUploadImageURL godoc
//
//	@Summary		Get presigned URL for image upload
//	@Description	Get a presigned URL to upload product images
//	@Tags			Product
//	@Produce		json
//	@Success		200	{object}	UploadImageURLResponseDto
//	@Failure		500	{object}	Error
//	@Router			/products/images/upload-url [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) GetUploadImageURL(ctx *gin.Context) {
	uploadURL, err := h.productApp.GetUploadImageURL(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, uploadURL)
}

// GetDeleteImageURL godoc
//
//	@Summary		Get presigned URL for image deletion
//	@Description	Get a presigned URL to delete product images
//	@Tags			Product
//	@Produce		json
//	@Param			image_id	path		string	true	"Product Image ID"	format(uuid)
//	@Success		200			{object}	DeleteImageURLResponseDto
//	@Failure		400			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/images/delete-url/{image_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) GetDeleteImageURL(ctx *gin.Context) {
	imageID, ok := pathToUUID(ctx, "image_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError("invalid image_id"))
		return
	}

	deleteURL, err := h.productApp.GetDeleteImageURL(ctx.Request.Context(), imageID)
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, deleteURL)
}
