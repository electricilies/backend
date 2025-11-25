package http

import (
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	List(*gin.Context)
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}
