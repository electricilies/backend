package error

import (
	domainerror "backend/internal/domain/error"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	var notFoundErr *domainerror.NotFoundError
	var conflictErr *domainerror.ConflictError
	var validationErr *domainerror.ValidationError
	var badRequestErr *domainerror.BadRequestError
	var internalErr *domainerror.InternalError
	var connectionErr *domainerror.ConnectionError
	var unavailableErr *domainerror.UnavailableError

	var code string
	var message string

	var domainErr *domainerror.DomainError
	if errors.As(err, &domainErr) {
		code = domainErr.Code
		message = domainErr.Message
	} else {
		message = err.Error()
	}

	switch {
	case errors.As(err, &notFoundErr):
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": message,
			"code":  code,
		})
	case errors.As(err, &validationErr):
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
			"code":  code,
		})
	case errors.As(err, &badRequestErr):
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
			"code":  code,
		})
	case errors.As(err, &conflictErr):
		ctx.JSON(http.StatusConflict, gin.H{
			"error": message,
			"code":  code,
		})
	case errors.As(err, &connectionErr), errors.As(err, &unavailableErr):
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": message,
			"code":  code,
		})
	case errors.As(err, &internalErr):
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": message,
			"code":  code,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}
}
