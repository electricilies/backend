package http

import (
	"net/http"
	"strings"

	"backend/config"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	cfgSrv *config.Server
}

func ProvideAuthHandler(cfg *config.Server) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		cfgSrv: cfg,
	}
}

func (h *AuthHandlerImpl) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		basePath := h.cfgSrv.KCBasePath
		redirectURL := basePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
