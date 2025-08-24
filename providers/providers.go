package providers

import (
	exProvider "project-root/modules/examples/providers"

	"gorm.io/gorm"
)

type Providers struct {
	Examples *exProvider.Provider
}

func Init(db *gorm.DB) *Providers {
	return &Providers{
		Examples: exProvider.NewProvider(db),
	}
}
