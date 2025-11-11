package handler

import (
	"net/http"

	"backend/internal/application"
	"backend/internal/interface/api/request"
	"backend/internal/interface/api/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetReturnRequests(ctx *gin.Context)
}

type userHandler struct {
	app application.User
}

func NewUser(app application.User) User {
	return &userHandler{app: app}
}

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{object}	User
//	@Failure		400		{object}	response.BadRequestError		"bad request"
//	@Failure		404		{object}	response.NotFoundError			"not found"
//	@Failure		500		{object}	response.InternalServerError	"internal error"
//	@Router			/users/{user_id} [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *userHandler) Get(ctx *gin.Context) {
	id := ctx.Param("user_id")
	u, err := h.app.Get(ctx, id)
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// ListUsers godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		response.User
//	@Failure		500	{object}	response.InternalServerError
//	@Router			/users [get]
//
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *userHandler) List(ctx *gin.Context) {
	users, err := h.app.List(ctx)
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.UsersFromDomain(users))
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		request.CreateUser	true	"User request"
//	@Success		201		{object}	response.User
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		409		{object}	response.ConflictError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/users [post]
func (h *userHandler) Create(ctx *gin.Context) {
	var req request.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendBadRequestError(ctx, err.Error())
		return
	}

	u := req.ToDomain()
	created, err := h.app.Create(ctx, u)
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string				true	"User ID"
//	@Param			user	body		request.UpdateUser	true	"User request"
//	@Success		204		{string}	string				"no content"
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		404		{object}	response.NotFoundError
//	@Failure		409		{object}	response.ConflictError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/users/{user_id} [put]
func (h *userHandler) Update(ctx *gin.Context) {
	id := ctx.Param("user_id")

	var req request.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := req.ToDomain()
	u.ID = uuid.MustParse(id)

	if err := h.app.Update(ctx, u); err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		204		{string}	string	"no content"
//	@Failure		404		{object}	response.NotFoundError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/users/{user_id} [delete]
func (h *userHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("user_id")
	if err := h.app.Delete(ctx, id); err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GetReturnRequests godoc
//
//	@Summary		Get return requests for a user
//	@Description	Get return requests for a user by user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{array}		response.ReturnRequest
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/users/{user_id}/returns [get]
func (h *userHandler) GetReturnRequests(ctx *gin.Context) {
	// TODO: implement getting return requests for a user
	ctx.Status(http.StatusNoContent)
}
