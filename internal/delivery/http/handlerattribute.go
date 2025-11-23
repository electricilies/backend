package http

import (
	"github.com/gin-gonic/gin"
)

type AttributeHandler interface {
	Get(*gin.Context)
	ListValues(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	CreateValue(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	DeleteValue(*gin.Context)
	UpdateValue(*gin.Context)
}
