package providers

import (
	"project-root/modules/examples/handler"
	"project-root/modules/examples/repository"
	"project-root/modules/examples/usecase"

	"gorm.io/gorm"
)

type Provider struct {
	ExHandler *handler.ExampleHandler
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewExampleRepository(db)
	usecase := usecase.NewExampleUsecase(repo)
	handler := handler.NewExampleHandler(usecase)

	return &Provider{
		ExHandler: handler,
	}
}
