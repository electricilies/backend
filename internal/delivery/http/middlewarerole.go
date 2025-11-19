package http

import (
	"github.com/gin-gonic/gin"
)

type UserRole string

type RoleMiddleware interface {
	Handler([]UserRole) gin.HandlerFunc
}
