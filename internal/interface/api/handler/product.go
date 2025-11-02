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
