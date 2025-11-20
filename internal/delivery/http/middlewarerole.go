package http

import (
	"backend/internal/application"
	"github.com/gin-gonic/gin"
)

type RoleMiddleware interface {
	Handler([]application.UserRole) gin.HandlerFunc
}
