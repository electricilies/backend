package http

import (
	"net/http"
	"strings"

	"backend/config"

	"github.com/gin-gonic/gin"
)

type GinAuthHandler struct {
	cfgSrv *config.Server
}

func ProvideAuthHandler(cfg *config.Server) *GinAuthHandler {
	return &GinAuthHandler{
		cfgSrv: cfg,
	}
}

func (h *GinAuthHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		basePath := h.cfgSrv.KCBasePath
		redirectURL := basePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
