package handler

import (
	app "backend/internal/application"
	"backend/internal/interface/api/request"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userHandler struct {
	app app.User
}

func NewUserHandler(app app.User) User {
	return &userHandler{app: app}
}

//	@BasePath	/api

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	user.User
//	@Failure		400	{string}	string	"bad request"
//	@Failure		404	{string}	string	"not found"
//	@Failure		500	{string}	string	"internal error"
//	@Router			/users/{id} [get]
func (h *userHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	u, err := h.app.Get(ctx, id)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
//	@Success		200	{array}		user.User
//	@Failure		500	{string}	string	"internal error"
//	@Router			/users [get]
func (h *userHandler) List(ctx *gin.Context) {
	users, err := h.app.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		request.User	true	"User request"
//	@Success		201		{object}	user.User
//	@Failure		400		{string}	string	"bad request"
//	@Failure		500		{string}	string	"internal error"
//	@Router			/users [post]
func (h *userHandler) Create(ctx *gin.Context) {
	var req request.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := req.ToDomain()
	created, err := h.app.Create(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
//	@Param			id		path		string			true	"User ID"
//	@Param			user	body		request.User	true	"User request"
//	@Success		204		{string}	string			"no content"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		404		{string}	string			"not found"
//	@Failure		500		{string}	string			"internal error"
//	@Router			/users/{id} [put]
func (h *userHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var req request.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := req.ToDomain()
	u.ID = id

	if err := h.app.Update(ctx, u); err != nil {
		if errors.Is(err, app.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
//	@Param			id	path		string	true	"User ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{string}	string	"not found"
//	@Failure		500	{string}	string	"internal error"
//	@Router			/users/{id} [delete]
func (h *userHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.app.Delete(ctx, id); err != nil {
		if errors.Is(err, app.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
