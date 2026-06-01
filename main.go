package main

import (
	"fmt"
	"os"
	"project-root/config"
	"project-root/internal/services"
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.InitEnv()

	db := config.InitDB()
	redis := config.InitRedis()
	jwtService := services.NewJWTService(
		os.Getenv("ACCESS_SECRET_KEY"),
		os.Getenv("REFRESH_SECRET_KEY"),
	)

	// TODO: Implement RBAC to users
	db.AutoMigrate(&model.Example{})

	p := providers.Init(db, redis, jwtService)
	r := gin.Default()
	routes.InitRoutes(r, p, jwtService)

	port := 8000
	fmt.Printf("Server running at port %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
