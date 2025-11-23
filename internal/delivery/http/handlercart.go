package http

import (
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	Get(*gin.Context)
	CreateItem(*gin.Context)
	UpdateItem(*gin.Context)
	RemoveItem(*gin.Context)
}
