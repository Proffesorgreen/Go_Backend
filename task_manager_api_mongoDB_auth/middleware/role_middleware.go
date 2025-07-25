package middleware

import "github.com/gin-gonic/gin"

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.IndentedJSON(403, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		if userRole != requiredRole {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden – insufficient permissions"})
			return
		}

		c.Next()
	}
}
