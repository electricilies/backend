package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrCodeNotFound     = "NOT_FOUND_ERROR"
	ErrCodeBadRequest   = "BAD_REQUEST_ERROR"
	ErrCodeConflict     = "CONFLICT_ERROR"
	ErrCodeUnavailable  = "SERVICE_UNAVAILABLE_ERROR"
	ErrCodeInternal     = "INTERNAL_ERROR"
	ErrCodeUnauthorized = "UNAUTHORIZED_ERROR"
	ErrCodeFobidden     = "FORBIDDEN_ERROR"
)

type NotFoundError struct {
	Error string `json:"error" example:"User with ID 123 not found"`
	Code  string `json:"code" example:"NOT_FOUND_ERROR"`
}

type BadRequestError struct {
	Error string `json:"error" example:"Email address is invalid"`
	Code  string `json:"code" example:"BAD_REQUEST_ERROR"`
}

type ConflictError struct {
	Error string `json:"error" example:"User with email already exists"`
	Code  string `json:"code" example:"CONFLICT_ERROR"`
}

type ServiceUnavailableError struct {
	Error string `json:"error" example:"Database connection failed"`
	Code  string `json:"code" example:"SERVICE_UNAVAILABLE_ERROR"`
}

type InternalServerError struct {
	Error string `json:"error" example:"An unexpected error occurred"`
	Code  string `json:"code" example:"INTERNAL_ERROR"`
}

type UnauthorizedError struct {
	Error string `json:"error" example:"Unauthorized access"`
	Code  string `json:"code" example:"UNAUTHORIZED_ERROR"`
}

type ForbiddenError struct {
	Error string `json:"error" example:"Forbidden access"`
	Code  string `json:"code" example:"FORBIDDEN_ERROR"`
}

// TODO: implement error handling from domain errors
func ErrorFromDomain(ctx *gin.Context, err error) {
	// if err == nil {
	// 	return
	// }
	//
	// var notFoundErr *domain.NotFoundError
	// var conflictErr *domain.ConflictError
	// var validationErr *domain.ValidationError
	// var badRequestErr *domain.BadRequestError
	// var internalErr *domain.InternalError
	// var connectionErr *domain.ConnectionError
	// var unavailableErr *domain.UnavailableError
	// var unauthorizedErr *domain.UnauthorizedError
	// var forbiddenErr *domain.ForbiddenError
	//
	// message := err.Error()
	//
	// switch {
	// case errors.As(err, &notFoundErr):
	// 	SendNotFoundError(ctx, message)
	// case errors.As(err, &validationErr), errors.As(err, &badRequestErr):
	// 	SendBadRequestError(ctx, message)
	// case errors.As(err, &conflictErr):
	// 	SendConflictError(ctx, message)
	// case errors.As(err, &connectionErr), errors.As(err, &unavailableErr):
	// 	SendServiceUnavailableError(ctx, message)
	// case errors.As(err, &internalErr):
	// 	SendInternalServerError(ctx, message)
	// case errors.As(err, &unauthorizedErr):
	// 	SendUnauthorizedError(ctx, message)
	// case errors.As(err, &forbiddenErr):
	// 	SendForbiddenError(ctx, message)
	//
	// default:
	// 	SendInternalServerError(ctx, "Internal server error")
	// }
}

func SendNotFoundError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, NotFoundError{
		Error: message,
		Code:  ErrCodeNotFound,
	})
}

func SendBadRequestError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, BadRequestError{
		Error: message,
		Code:  ErrCodeBadRequest,
	})
}

func SendConflictError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusConflict, ConflictError{
		Error: message,
		Code:  ErrCodeConflict,
	})
}

func SendServiceUnavailableError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusServiceUnavailable, ServiceUnavailableError{
		Error: message,
		Code:  ErrCodeUnavailable,
	})
}

func SendInternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, InternalServerError{
		Error: message,
		Code:  ErrCodeInternal,
	})
}

func SendUnauthorizedError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized,
		ForbiddenError{
			Error: message,
			Code:  ErrCodeUnauthorized,
		})
}

func SendForbiddenError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusForbidden,
		ForbiddenError{
			Error: message,
			Code:  ErrCodeFobidden,
		})
}
