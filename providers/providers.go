package providers

import (
	authProvider "project-root/modules/auth/providers"
	exProvider "project-root/modules/examples/providers"
	userProvider "project-root/modules/users/providers"

	"gorm.io/gorm"
)

type Providers struct {
	Examples *exProvider.Provider
	Users    *userProvider.Provider
	Auth     *authProvider.Provider
}

func Init(db *gorm.DB) *Providers {
	return &Providers{
		Examples: exProvider.NewProvider(db),
		Users:    userProvider.NewProvider(db),
		Auth:     authProvider.NewProvider(db),
	}
}
