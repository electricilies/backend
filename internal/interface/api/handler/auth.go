package handler

import (
	"net/http"
	"strings"

	"backend/config"

	"github.com/gin-gonic/gin"
)

type Auth interface {
	Handler() gin.HandlerFunc
}

type AuthImpl struct {
	cfgSrv *config.Server
}

func NewAuth(cfgSrv *config.Server) Auth {
	return &AuthImpl{
		cfgSrv: cfgSrv,
	}
}

func ProvideAuth(cfg *config.Server) *AuthImpl {
	return &AuthImpl{
		cfgSrv: cfg,
	}
}

func (h *AuthImpl) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		basePath := h.cfgSrv.KCBasePath
		redirectURL := basePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
