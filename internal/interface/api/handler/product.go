package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	ListVariantsByProduct(ctx *gin.Context)
	ListReviewByProduct(ctx *gin.Context)
	ListAtributesByProduct(ctx *gin.Context)
	AddAttributeValues(ctx *gin.Context)
	RemoveAttributeValue(ctx *gin.Context)
}

type productHandler struct{}

func NewProduct() Product { return &productHandler{} }

// GetProduct godoc
//
//	@Summary		Get product by ID
//	@Description	Get product details by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	response.Product
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products/{id} [get]
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
//	@Success		201		{object}	response.Product
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
//	@Param			id		path		int						true	"Product ID"
//	@Param			product	body		request.UpdateProduct	true	"Update product request"
//	@Success		204		{string}	string					"no content"
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		404		{object}	mapper.NotFoundError
//	@Failure		409		{object}	mapper.ConflictError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/products/{id} [put]
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
//	@Param			id	path		int		true	"Product ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products/{id} [delete]
func (h *productHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListVariantsByProduct godoc
//
//	@Summary		List variants by product
//	@Description	Get all variants for a specific product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{array}		response.ProductVariant
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products/{id}/variants [get]
func (h *productHandler) ListVariantsByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListReviewByProduct godoc
//
//	@Summary		List reviews by product
//	@Description	Get all reviews for a specific product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{array}		response.Review
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products/{id}/reviews [get]
func (h *productHandler) ListReviewByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListAtributesByProduct godoc
//
//	@Summary		List attributes by product
//	@Description	Get all attributes for a specific product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{array}		response.Attribute
//	@Failure		404	{object}	mapper.NotFoundError
//	@Failure		500	{object}	mapper.InternalServerError
//	@Router			/products/{id}/attributes [get]
func (h *productHandler) ListAtributesByProduct(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// AddAttributeValues godoc
//
//	@Summary		Add attribute values to product
//	@Description	Add attribute values to a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Product ID"
//	@Param			values	body		request.AddAttributeValues	true	"Add attribute values request"
//	@Success		204		{string}	string						"no content"
//	@Failure		400		{object}	mapper.BadRequestError
//	@Failure		404		{object}	mapper.NotFoundError
//	@Failure		500		{object}	mapper.InternalServerError
//	@Router			/products/{id}/attributes [post]
func (h *productHandler) AddAttributeValues(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// RemoveAttributeValue godoc
//
//	@Summary		Remove attribute value from product
//	@Description	Remove attribute value from a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int		true	"Product ID"
//	@Param			value_id	path		int		true	"Attribute Value ID"
//	@Success		204			{string}	string	"no content"
//	@Failure		404			{object}	mapper.NotFoundError
//	@Failure		500			{object}	mapper.InternalServerError
//	@Router			/products/{id}/attributes/{value_id} [delete]
func (h *productHandler) RemoveAttributeValue(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
