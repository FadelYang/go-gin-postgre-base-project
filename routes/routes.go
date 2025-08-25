package routes

import (
	"project-root/providers"

	"github.com/gin-gonic/gin"

	ex "project-root/modules/examples/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine, p *providers.Providers) {
	api := r.Group("api/v1")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ex.RegisterRoutes(api, p.Examples)
}
