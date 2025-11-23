package http

import (
	"net/http"

	_ "backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	AddImages(*gin.Context)
	DeleteImages(*gin.Context)
	AddVariants(*gin.Context)
	UpdateVariant(*gin.Context)
	UpdateOptions(*gin.Context)
	GetDeleteImageURL(*gin.Context)
	GetUploadImageURL(*gin.Context)
}


