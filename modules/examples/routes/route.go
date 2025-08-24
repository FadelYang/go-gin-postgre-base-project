package routes

import (
	"project-root/modules/examples/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, exProvider *providers.Provider) {
	exRoutes := rg.Group("/examples")

	exRoutes.GET("", exProvider.ExController.GetExamples)
	exRoutes.POST("", exProvider.ExController.Create)
}
