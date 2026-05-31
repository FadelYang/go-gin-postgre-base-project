package providers

import (
	"project-root/internal/services"
	"project-root/modules/auth/handler"
	"project-root/modules/auth/repository"
	"project-root/modules/auth/usecase"
	userRepository "project-root/modules/users/repository"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Provider struct {
	AuthHandler *handler.AuthHandler
}

func NewProvider(db *gorm.DB, redisClient *redis.Client, jwtService *services.JWTService) *Provider {
	repo := repository.NewAuthRepository(db)
	userRepo := userRepository.NewuserRepository(db)
	usecase := usecase.NewAuthUsecase(redisClient, repo, userRepo, jwtService)
	handler := handler.NewAuthHandler(usecase)

	return &Provider{
		AuthHandler: handler,
	}
}
