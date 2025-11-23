package http

import (
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	List(*gin.Context)
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}
