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

// Get godoc
//
//	@Summary		Health check
//	@Description	Returns current server time to verify service is running
//	@Tags			Health
//	@Produce		json
//	@Success		200	{string}	string	"current server time"
//	@Router			/health [get]
func (h *healthCheck) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now())
}
