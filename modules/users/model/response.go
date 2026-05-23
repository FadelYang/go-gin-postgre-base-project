package model

import (
	"time"

	"github.com/google/uuid"
)

type UserDetail struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Phonenumber  string    `json:"phonenumber"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PasswordHash string    `json:"password_hash"`
	TokenVersion int       `json:"token_version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
