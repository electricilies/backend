package mapper

import (
	"backend/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundError struct {
	Error string `json:"error" example:"User with ID 123 not found"`
	Code  string `json:"code" example:"USER_NOT_FOUND"`
}

type BadRequestError struct {
	Error string `json:"error" example:"Email address is invalid"`
	Code  string `json:"code" example:"INVALID_EMAIL"`
}

type ConflictError struct {
	Error string `json:"error" example:"User with email already exists"`
	Code  string `json:"code" example:"EMAIL_EXISTS"`
}

type ServiceUnavailableError struct {
	Error string `json:"error" example:"Database connection failed"`
	Code  string `json:"code" example:"DB_UNAVAILABLE"`
}

type InternalServerError struct {
	Error string `json:"error" example:"An unexpected error occurred"`
	Code  string `json:"code" example:"INTERNAL_ERROR"`
}

func ErrorFromDomain(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	var notFoundErr *domain.NotFoundError
	var conflictErr *domain.ConflictError
	var validationErr *domain.ValidationError
	var badRequestErr *domain.BadRequestError
	var internalErr *domain.InternalError
	var connectionErr *domain.ConnectionError
	var unavailableErr *domain.UnavailableError

	var code string
	var message string

	var domainErr *domain.Err
	if errors.As(err, &domainErr) {
		code = domainErr.Code
		message = domainErr.Message
	} else {
		message = err.Error()
	}

	switch {
	case errors.As(err, &notFoundErr):
		SendNotFoundError(ctx, message, code)
	case errors.As(err, &validationErr), errors.As(err, &badRequestErr):
		SendBadRequestError(ctx, message, code)
	case errors.As(err, &conflictErr):
		SendConflictError(ctx, message, code)
	case errors.As(err, &connectionErr), errors.As(err, &unavailableErr):
		SendServiceUnavailableError(ctx, message, code)
	case errors.As(err, &internalErr):
		SendInternalServerError(ctx, message, code)
	default:
		SendInternalServerError(ctx, "Internal server error", "")
	}
}

func SendNotFoundError(ctx *gin.Context, message string, code string) {
	ctx.JSON(http.StatusNotFound, NotFoundError{
		Error: message,
		Code:  code,
	})
}

func SendBadRequestError(ctx *gin.Context, message string, code string) {
	ctx.JSON(http.StatusBadRequest, BadRequestError{
		Error: message,
		Code:  code,
	})
}

func SendConflictError(ctx *gin.Context, message string, code string) {
	ctx.JSON(http.StatusConflict, ConflictError{
		Error: message,
		Code:  code,
	})
}

func SendServiceUnavailableError(ctx *gin.Context, message string, code string) {
	ctx.JSON(http.StatusServiceUnavailable, ServiceUnavailableError{
		Error: message,
		Code:  code,
	})
}

func SendInternalServerError(ctx *gin.Context, message string, code string) {
	ctx.JSON(http.StatusInternalServerError, InternalServerError{
		Error: message,
		Code:  code,
	})
}
