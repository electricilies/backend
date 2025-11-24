package http

import (
	"fmt"
	"net/http"

	"backend/internal/application"
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandlerImpl struct {
	productApp           application.Product
	ErrRequiredProductID string
	ErrInvalidProductID  string
}

var _ ProductHandler = &ProductHandlerImpl{}

func ProvideProductHandler(productApp application.Product) *ProductHandlerImpl {
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
//	@Success		200			{object}	domain.Product
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id} [get]
func (h *ProductHandlerImpl) Get(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}
	product, err := h.productApp.Get(ctx.Request.Context(), application.GetProductParam{
		ProductID: *productID,
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
//	@Success		200				{object}	application.Pagination[domain.Product]
//	@Failure		500				{object}	Error
//	@Router			/products [get]
func (h *ProductHandlerImpl) List(ctx *gin.Context) {
	paginateParam, err := createPaginationParamsFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	var productIDs *[]uuid.UUID
	if productIDsQuery, ok := queryArrayToUUIDSlice(ctx, "product_ids"); ok {
		productIDs = productIDsQuery
	}

	var categoryIDs *[]uuid.UUID
	if categoryIDsQuery, ok := queryArrayToUUIDSlice(ctx, "category_ids"); ok {
		categoryIDs = categoryIDsQuery
	}

	var search *string
	if searchQuery, ok := ctx.GetQuery("search"); ok {
		search = &searchQuery
	}

	var minPrice *int64
	if minPriceQuery, ok := ctx.GetQuery("min_price"); ok {
		var price int64
		if _, err := fmt.Sscanf(minPriceQuery, "%d", &price); err == nil {
			minPrice = &price
		}
	}

	var maxPrice *int64
	if maxPriceQuery, ok := ctx.GetQuery("max_price"); ok {
		var price int64
		if _, err := fmt.Sscanf(maxPriceQuery, "%d", &price); err == nil {
			maxPrice = &price
		}
	}

	var rating *float64
	if ratingQuery, ok := ctx.GetQuery("rating"); ok {
		var r float64
		if _, err := fmt.Sscanf(ratingQuery, "%f", &r); err == nil {
			rating = &r
		}
	}

	var sortPrice *string
	if sortPriceQuery, ok := ctx.GetQuery("sort_price"); ok {
		sortPrice = &sortPriceQuery
	}

	var sortRating *string
	if sortRatingQuery, ok := ctx.GetQuery("sort_rating"); ok {
		sortRating = &sortRatingQuery
	}

	deleted := domain.DeletedExcludeParam
	if deletedQuery, ok := ctx.GetQuery("deleted"); ok {
		deleted = domain.DeletedParam(deletedQuery)
	}

	products, err := h.productApp.List(ctx.Request.Context(), application.ListProductParam{
		PaginationParam: *paginateParam,
		ProductIDs:      productIDs,
		CategoryIDs:     categoryIDs,
		MinPrice:        minPrice,
		MaxPrice:        maxPrice,
		Rating:          rating,
		SortPrice:       sortPrice,
		SortRating:      sortRating,
		Search:          search,
		Deleted:         deleted,
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
//	@Param			product	body		application.CreateProductData	true	"Product request"
//	@Success		201		{object}	domain.Product
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/products [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) Create(ctx *gin.Context) {
	var data application.CreateProductData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product, err := h.productApp.Create(ctx.Request.Context(), application.CreateProductParam{
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
//	@Param			product_id	path		string							true	"Product ID"	format(uuid)
//	@Param			product		body		application.UpdateProductData	true	"Update product request"
//	@Success		200			{object}	domain.Product
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) Update(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data application.UpdateProductData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product, err := h.productApp.Update(ctx.Request.Context(), application.UpdateProductParam{
		ProductID: *productID,
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
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	err := h.productApp.Delete(ctx.Request.Context(), application.DeleteProductParam{
		ProductID: *productID,
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
//	@Param			product_id		path		string								true	"Product ID"	format(uuid)
//	@Param			productImages	body		[]application.AddProductImageData	true	"Product images request"
//	@Success		201				{array}		domain.ProductImage
//	@Failure		400				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/images [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) AddImages(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data []application.AddProductImageData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	images, err := h.productApp.AddImages(ctx.Request.Context(), application.AddProductImagesParam{
		ProductID: *productID,
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
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	imageIDs, ok := queryArrayToUUIDSlice(ctx, "ids")
	if !ok || imageIDs == nil || len(*imageIDs) == 0 {
		ctx.JSON(http.StatusBadRequest, NewError("image ids are required"))
		return
	}

	err := h.productApp.DeleteImages(ctx.Request.Context(), application.DeleteProductImagesParam{
		ProductID: *productID,
		ImageIDs:  *imageIDs,
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
//	@Param			product_id		path		string									true	"Product ID"	format(uuid)
//	@Param			productVariant	body		[]application.AddProductVariantsData	true	"Product variant request"
//	@Success		201				{array}		domain.ProductVariant
//	@Failure		400				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/variants [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) AddVariants(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data []application.AddProductVariantsData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	variants, err := h.productApp.AddVariants(ctx.Request.Context(), application.AddProductVariantsParam{
		ProductID: *productID,
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
//	@Param			product_id		path		string									true	"Product ID"			format(uuid)
//	@Param			variant_id		path		string									true	"Product Variant ID"	format(uuid)
//	@Param			productVariant	body		application.UpdateProductVariantData	true	"Update product variant request"
//	@Success		200				{object}	domain.ProductVariant
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		409				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/products/{product_id}/variants/{variant_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) UpdateVariant(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	variantID, ok := pathToUUID(ctx, "variant_id")
	if *variantID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError("variant_id is required"))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError("invalid variant_id"))
		return
	}

	var data application.UpdateProductVariantData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	variant, err := h.productApp.UpdateVariant(ctx.Request.Context(), application.UpdateProductVariantParam{
		ProductID:        *productID,
		ProductVariantID: *variantID,
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
//	@Param			product_id	path		string									true	"Product ID"	format(uuid)
//	@Param			option		body		[]application.UpdateProductOptionsData	true	"Update product option request"
//	@Success		200			{array}		domain.Option
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/{product_id}/options [put]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) UpdateOptions(ctx *gin.Context) {
	productID, ok := pathToUUID(ctx, "product_id")
	if *productID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredProductID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidProductID))
		return
	}

	var data []application.UpdateProductOptionsData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	options, err := h.productApp.UpdateOptions(ctx.Request.Context(), application.UpdateProductOptionsParam{
		ProductID: *productID,
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
//	@Success		200	{object}	application.UploadImageURL
//	@Failure		500	{object}	Error
//	@Router			/products/images/upload-url [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) GetUploadImageURL(ctx *gin.Context) {
	uploadURL, err := h.productApp.GetUploadImageURL(ctx.Request.Context())
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
//	@Success		200			{object}	application.DeleteImageURL
//	@Failure		400			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/products/images/delete-url/{image_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *ProductHandlerImpl) GetDeleteImageURL(ctx *gin.Context) {
	imageID, ok := pathToUUID(ctx, "image_id")
	if *imageID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError("image_id is required"))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError("invalid image_id"))
		return
	}

	deleteURL, err := h.productApp.GetDeleteImageURL(ctx.Request.Context(), *imageID)
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, deleteURL)
}
