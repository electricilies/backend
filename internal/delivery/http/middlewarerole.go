package http

import (
	"backend/config"

	"github.com/gin-gonic/gin"
)

type UserRole string

type RoleMiddleware interface {
	Handler([]UserRole) gin.HandlerFunc
}

type RoleMiddlewareImpl struct {
	srvCfg *config.Server
}
