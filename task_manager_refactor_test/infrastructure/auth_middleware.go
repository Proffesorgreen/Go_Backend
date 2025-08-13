package infrastructure

import (
	"net/http"
	"strings"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtservice domain.JWT_Service
}

func NewAuthMiddleware(js domain.JWT_Service) *AuthMiddleware {
	return &AuthMiddleware{
		jwtservice: js,
	}
}

func (am *AuthMiddleware) Authentcate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "authorization header needed"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := am.jwtservice.ValidateJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("user_role", claims.Role)
		c.Set("username", claims.Username)

		c.Next()
	}
}

func (am *AuthMiddleware) AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "authentication context missing"})
			c.Abort()
			return
		}

		userRole, err := role.(domain.Role)
		if !err {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user format"})
			c.Abort()
			return
		}

		if userRole != domain.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "access forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
