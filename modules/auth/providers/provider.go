package providers

import (
	"project-root/modules/auth/handler"
	"project-root/modules/auth/repository"
	"project-root/modules/auth/usecase"
	userRepository "project-root/modules/users/repository"

	"gorm.io/gorm"
)

type Provider struct {
	AuthHandler *handler.AuthHandler
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewAuthRepository(db)
	userRepo := userRepository.NewuserRepository(db)
	usecase := usecase.NewAuthUsecase(repo, userRepo)
	handler := handler.NewAuthHandler(usecase)

	return &Provider{
		AuthHandler: handler,
	}
}
