package routes

import (
	"project-root/providers"

	"github.com/gin-gonic/gin"

	ex "project-root/modules/examples/routes"
)

func InitRoutes(r *gin.Engine, p *providers.Providers) {
	api := r.Group("api/v1")

	ex.RegisterRoutes(api, p.Examples)
}
