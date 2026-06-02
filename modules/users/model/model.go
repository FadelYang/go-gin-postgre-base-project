package model

import (
	roleModel "project-root/modules/roles/model"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"not null;unique"`
	Phonenumber  string    `gorm:"not null;unique"`
	Email        string    `gorm:"not null;unique"`
	FirstName    string    `gorm:"not null"`
	LastName     string    `gorm:""`
	PasswordHash string    `gorm:"not null"`
	TokenVersion int       `gorm:""`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	RoleID uuid.UUID       `gorm:"column:role_id"`
	Role   roleModel.Roles `gorm:"foreignKey:RoleID;references:ID"`
}
