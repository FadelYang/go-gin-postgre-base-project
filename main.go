package main

import (
	"fmt"
	"project-root/config"
	"project-root/modules/examples/model"
	"project-root/providers"
	"project-root/routes"

	"github.com/gin-gonic/gin"

	_ "project-root/docs"
)

// @title					Go, Gin, and Postgre Base Project
// @version				1.0
// @description 	Go, Gin, and Postgre Base Project
// @BasePath			/api/v1
func main() {
	config.InitEnv()

	db := config.InitDB()

	db.AutoMigrate(&model.Example{})

	p := providers.Init(db)
	r := gin.Default()
	routes.InitRoutes(r, p)

	port := 8000
	fmt.Printf("Server running at port %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
