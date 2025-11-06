package middleware

import (
	"net/http"

	"backend/config"
	"backend/internal/constant"

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

		if !roleAllowed(constant.UserRole(userRole), rolesAllowed) {
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

func roleAllowed(userRole constant.UserRole, allowedRole []constant.UserRole) bool {
	set := make(map[constant.UserRole]struct{})
	for _, requiredRole := range allowedRole {
		set[requiredRole] = struct{}{}
	}
	if _, exists := set[userRole]; exists {
		return true
	}
	return false
}
