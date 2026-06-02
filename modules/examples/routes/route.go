package routes

import (
	"project-root/internal/services"
	"project-root/middleware"
	"project-root/modules/examples/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, exProvider *providers.Provider, jwtService *services.JWTService) {
	exRoutes := rg.Group("/examples")

	exRoutes.GET("", exProvider.ExHandler.GetExamples)
	exRoutes.GET("/auth", middleware.AuthMiddleware(*jwtService), exProvider.ExHandler.GetExampleWithAuth)
	exRoutes.GET(
		"/admin-superadmin-only",
		middleware.AuthMiddleware(*jwtService),
		middleware.RBACMiddleware(*jwtService, []string{"superadmin", "admin"}),
		exProvider.ExHandler.GetExampleOnlyForAdminAndSuperAdmin,
	)
	exRoutes.POST("", exProvider.ExHandler.Create)
}
