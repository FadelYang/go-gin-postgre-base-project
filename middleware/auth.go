package middleware

import (
	"fmt"
	"net/http"
	"project-root/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing authorization header",
			})
			return
		}

		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization format",
			})
			return
		}

		accessToken := splitToken[1]

		claims, err := jwtService.ValidateAccessToken(accessToken)
		if err != nil {
			fmt.Printf("err validate token: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
