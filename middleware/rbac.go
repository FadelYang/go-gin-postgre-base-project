package middleware

import (
	"net/http"
	"project-root/internal/services"
	"project-root/modules/auth/dto"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(jwtService services.JWTService, roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsValue, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "claims not found",
			})
			return
		}

		claims, ok := claimsValue.(*dto.AccessTokenClaim)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid claims",
			})
			return
		}

		passed := false

		for _, v := range roles {
			if v == claims.Role {
				passed = true
			}
		}

		if !passed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "access denied",
			})
			return
		}

		c.Next()
	}
}
