package routes

import (
	"project-root/modules/auth/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, exProvider *providers.Provider) {
	exRoutes := rg.Group("/auth")

	exRoutes.POST("/register", exProvider.AuthHandler.Register)
}
