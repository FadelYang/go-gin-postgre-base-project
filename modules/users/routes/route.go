package routes

import (
	"project-root/modules/users/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, userProvider *providers.Provider) {
	exRoutes := rg.Group("/users")

	exRoutes.GET("", userProvider.UserController.GetAll)
	exRoutes.POST("", userProvider.UserController.Create)
	exRoutes.PUT(":uuid", userProvider.UserController.Update)
	exRoutes.DELETE("/:uuid", userProvider.UserController.Delete)
	exRoutes.GET("/:uuid", userProvider.UserController.GetByID)
	exRoutes.GET("/email/:email", userProvider.UserController.GetByEmail)
}
