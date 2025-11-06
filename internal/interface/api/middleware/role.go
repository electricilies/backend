package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleAdmin    UserRole = "admin"
	RoleStaff    UserRole = "staff"
)

type Role interface {
	Handler() gin.HandlerFunc
}

type roleMiddleware struct {
	clientId      string
	requiredRoles []UserRole
}

func NewRole(clientId string, requiredRoles []UserRole) Role {
	return &roleMiddleware{
		clientId:      clientId,
		requiredRoles: requiredRoles,
	}
}

func (r *roleMiddleware) Handler() gin.HandlerFunc {
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

		userRoles := extractRoles(claims, r.clientId)

		userRoleEnums := make([]UserRole, 0, len(userRoles))
		for _, role := range userRoles {
			userRoleEnums = append(userRoleEnums, UserRole(role))
		}

		if !hasRequiredRole(userRoleEnums, r.requiredRoles) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		ctx.Next()
	}
}

func extractRoles(claims jwt.MapClaims, clientID string) []string {
	var roles []string

	resAccess, _ := claims["resource_access"].(map[string]interface{})
	if resAccess == nil {
		return roles
	}

	clientRes, _ := resAccess[clientID].(map[string]interface{})
	if clientRes == nil {
		return roles
	}

	resRoles, _ := clientRes["roles"].([]interface{})
	for _, r := range resRoles {
		if role, _ := r.(string); role != "" {
			roles = append(roles, role)
		}
	}

	return roles
}

func hasRequiredRole(userRoles []UserRole, requiredRoles []UserRole) bool {
	if len(requiredRoles) == 0 {
		return true
	}

	roleMap := make(map[UserRole]bool)
	for _, role := range userRoles {
		roleMap[role] = true
	}

	for _, requiredRole := range requiredRoles {
		if roleMap[requiredRole] {
			return true
		}
	}

	return false
}
