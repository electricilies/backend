package middleware

import (
	"net/http"
	"strings"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	Handler() gin.HandlerFunc
}

type AuthMiddleware struct {
	keycloakClient *gocloak.GoCloak
	cfg            *config.Config
}

func NewAuth(keycloakClient *gocloak.GoCloak, cfg *config.Config) Auth {
	return &AuthMiddleware{
		keycloakClient: keycloakClient,
		cfg:            cfg,
	}
}

func ProvideAuth(keycloakClient *gocloak.GoCloak, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		keycloakClient: keycloakClient,
		cfg:            cfg,
	}
}

func (j *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		token := parts[1]
		tokens, _, err := j.keycloakClient.DecodeAccessToken(ctx, token, j.cfg.KCRealm)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot decode access token", "detail": err.Error()})
			return
		}
		rptResult, err := j.keycloakClient.RetrospectToken(ctx, token, j.cfg.KCClientId, j.cfg.KCClientSecret, j.cfg.KCRealm)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to introspect token", "detail": err.Error()})
			return
		}
		if rptResult == nil || !*rptResult.Active {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Inactive or invalid token"})
			return
		}

		claims, _ := tokens.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
