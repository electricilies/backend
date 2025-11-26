package http

import (
	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	Get(*gin.Context)
	GetByUser(*gin.Context)
	Create(*gin.Context)
	CreateItem(*gin.Context)
	UpdateItem(*gin.Context)
	RemoveItem(*gin.Context)
}
