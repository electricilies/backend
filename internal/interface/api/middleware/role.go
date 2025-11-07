package middleware

import (
	"backend/config"
	"backend/internal/constant"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Role interface {
	Handler(requiredRoles []constant.UserRole) gin.HandlerFunc
}

type roleMiddleware struct {
	clientId string
}

func NewRole(requiredRoles []constant.UserRole) Role {
	return &roleMiddleware{
		clientId: config.Cfg.KcClientId,
	}
}

func (r *roleMiddleware) Handler(rolesAllowed []constant.UserRole) gin.HandlerFunc {
	set := make(map[constant.UserRole]struct{})
	for _, requiredRole := range rolesAllowed {
		set[requiredRole] = struct{}{}
	}
	return func(ctx *gin.Context) {
		claimsInterface, exists := ctx.Get("claims")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No JWT claims found, JWT middleware must be used before role middleware"})
			return
		}

		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT claims format"})
			return
		}

		userRole := extractRoleFromRoot(claims)
		if userRole == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No role found in JWT claims"})
			return
		}

		if _, exists := set[constant.UserRole(userRole)]; !exists {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		ctx.Next()
	}
}

func extractRoleFromRoot(claims jwt.MapClaims) string {
	role, _ := claims["role"].(string)
	return role
}
