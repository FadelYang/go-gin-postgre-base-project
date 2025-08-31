package providers

import (
	exProvider "project-root/modules/examples/providers"
	userProvider "project-root/modules/users/providers"

	"gorm.io/gorm"
)

type Providers struct {
	Examples *exProvider.Provider
	Users    *userProvider.Provider
}

func Init(db *gorm.DB) *Providers {
	return &Providers{
		Examples: exProvider.NewProvider(db),
		Users:    userProvider.NewProvider(db),
	}
}
