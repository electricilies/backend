package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthCheck interface {
	Get(ctx *gin.Context)
}

type healthCheck struct {
}

func NewHealthCheck() HealthCheck {
	return &healthCheck{}
}

func (h *healthCheck) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now())
}
