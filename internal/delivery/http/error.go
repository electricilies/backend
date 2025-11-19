package http

import (
	"errors"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func SendError(ctx *gin.Context, err error) {
	var httpErrCode int
	switch {
	case errors.Is(err, domain.ErrNotFound):
		httpErrCode = 404
	case errors.Is(err, domain.ErrInvalid), errors.Is(err, domain.ErrInvalidFormat):
		httpErrCode = 400
	case errors.Is(err, domain.ErrExists), errors.Is(err, domain.ErrConflict):
		httpErrCode = 409
	case errors.Is(err, domain.ErrForbidden):
		httpErrCode = 403
	case errors.Is(err, domain.ErrUnavailable):
		httpErrCode = 503
	case errors.Is(err, domain.ErrTimeout):
		httpErrCode = 504
	case errors.Is(err, domain.ErrConnection):
		httpErrCode = 502
	case errors.Is(err, domain.ErrCanceled):
		httpErrCode = 499
	default:
		httpErrCode = 500
	}
	ctx.JSON(httpErrCode, Error{Message: errors.Unwrap(err).Error()})
}
