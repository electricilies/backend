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

type AuthHandler struct {
	cfg *config.Config
}

func NewAuth(cfg *config.Config) Auth {
	return &AuthHandler{
		cfg: cfg,
	}
}

func ProvideAuth(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		basePath := h.cfg.KCBasePath
		redirectURL := basePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
