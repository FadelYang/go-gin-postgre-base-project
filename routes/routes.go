package routes

import (
	"project-root/providers"

	"github.com/gin-gonic/gin"

	ex "project-root/modules/examples/routes"
	users "project-root/modules/users/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine, p *providers.Providers) {
	api := r.Group("api/v1")

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ex.RegisterRoutes(api, p.Examples)
	users.RegisterRoutes(api, p.Users)
}
