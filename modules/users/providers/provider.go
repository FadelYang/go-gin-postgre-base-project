package providers

import (
	"project-root/modules/users/handler"
	"project-root/modules/users/repository"
	"project-root/modules/users/usecase"

	"gorm.io/gorm"
)

type Provider struct {
	UserHandler *handler.UserHandler
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewuserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(usecase)

	return &Provider{
		UserHandler: handler,
	}
}
