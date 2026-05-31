package providers

import (
	"project-root/internal/services"
	authProvider "project-root/modules/auth/providers"
	exProvider "project-root/modules/examples/providers"
	userProvider "project-root/modules/users/providers"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Providers struct {
	Examples *exProvider.Provider
	Users    *userProvider.Provider
	Auth     *authProvider.Provider
}

func Init(db *gorm.DB, redisClient *redis.Client, jwtService *services.JWTService) *Providers {
	return &Providers{
		Examples: exProvider.NewProvider(db),
		Users:    userProvider.NewProvider(db),
		Auth:     authProvider.NewProvider(db, redisClient, jwtService),
	}
}
