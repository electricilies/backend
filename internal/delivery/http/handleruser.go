package http

import (
	"net/http"

	_ "backend/internal/domain"
	_ "backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetCart(*gin.Context)
}

type GinUserHandler struct{}

var _ UserHandler = &GinUserHandler{}

func ProvideUserHandler() *GinUserHandler {
	return &GinUserHandler{}
}

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{object}	domain.User
//	@Failure		400		{object}	service.BadRequestError		"bad request"
//	@Failure		404		{object}	service.NotFoundError			"not found"
//	@Failure		500		{object}	service.InternalServerError	"internal error"
//	@Router			/users/{user_id} [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *GinUserHandler) Get(ctx *gin.Context) {
}

// ListUsers godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array} domain.User
//	@Failure		500	{object}	service.InternalServerError
//	@Router			/users [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *GinUserHandler) List(ctx *gin.Context) {
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body service.CreateUserParam	true	"User request"
//	@Success		201		{object} domain.User
//	@Failure		400		{object}	service.BadRequestError
//	@Failure		409		{object}	service.ConflictError
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/users [post]
func (h *GinUserHandler) Create(ctx *gin.Context) {
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string				true	"User ID"
//	@Param			user	body service.UpdateUserParam	true	"User request"
//	@Success		204		{object} domain.User
//	@Failure		400		{object}	service.BadRequestError
//	@Failure		404		{object}	service.NotFoundError
//	@Failure		409		{object}	service.ConflictError
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/users/{user_id} [patch]
func (h *GinUserHandler) Update(ctx *gin.Context) {
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		204
//	@Failure		404		{object}	service.NotFoundError
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/users/{user_id} [delete]
func (h *GinUserHandler) Delete(ctx *gin.Context) {
}

// GetCart godoc
//
//	@Summary		Get cart for a user
//	@Description	Get cart for a user by user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{object} domain.Cart
//	@Failure		400		{object}	service.BadRequestError
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/users/{user_id}/cart [get]
func (h *GinUserHandler) GetCart(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
