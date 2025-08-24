package providers

import (
	"project-root/modules/examples/controller"
	"project-root/modules/examples/repository"
	"project-root/modules/examples/service"

	"gorm.io/gorm"
)

type Provider struct {
	ExController *controller.ExampleController
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewExampleRepository(db)
	service := service.NewExampleService(repo)
	controller := controller.NewExampleController(service)

	return &Provider{
		ExController: controller,
	}
}
