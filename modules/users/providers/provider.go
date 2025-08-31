package providers

import (
	"project-root/modules/users/controller"
	"project-root/modules/users/repository"
	"project-root/modules/users/service"

	"gorm.io/gorm"
)

type Provider struct {
	UserController *controller.UserController
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewuserRepository(db)
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)

	return &Provider{
		UserController: controller,
	}
}
