package http

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler interface {
	Liveness(*gin.Context)
	Readiness(*gin.Context)
}
