package routes

import (
	"project-root/modules/users/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, userProvider *providers.Provider) {
	exRoutes := rg.Group("/users")

	exRoutes.GET("", userProvider.UserHandler.GetAll)
	exRoutes.POST("", userProvider.UserHandler.Create)
	exRoutes.PUT(":uuid", userProvider.UserHandler.Update)
	exRoutes.DELETE("/:uuid", userProvider.UserHandler.Delete)
	exRoutes.GET("/:uuid", userProvider.UserHandler.GetByID)
	exRoutes.GET("/email/:email", userProvider.UserHandler.GetByEmail)

	// TODO: add RBAC middleware
	// TODO: implement RBAC middleware that only user with superadmin role can update other's role
	exRoutes.PUT("/:uuid/role", userProvider.UserHandler.UpdateRole)
}
