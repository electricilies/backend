package http

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
