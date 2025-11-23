package http

import (
	"net/http"

	"backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}


