package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	UUID         uuid.UUID `json:"id" example:"a53515e3-5a7f-440b-82f6-3d84ac7ce746"`
	Username     string    `json:"username" example:"Budi Pambudi"`
	Email        string    `json:"email" example:"budipambudi@gmail.com"`
	PasswordHash string    `json:"-" example:"$2a$10$N9qo8uLOickgx2ZMRZo5i.eW5qk8Qy4T/3yFz5eF5r5eF5r5eF5r5e"`
	CreatedAt    time.Time `json:"created_at" example:"1617181723"`
	UpdatedAt    time.Time `json:"updated_at" example:"1617181723"`
}

type CreateUser struct {
	Username     string    `json:"username" example:"Budi Pambudi"`
	Email        string    `json:"email" example:"budipambudi@gmail.com"`
	Password     string    `json:"password" example:"supersecretpassword"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type UpdateUser struct {
	Username string `json:"username" example:"Budi Pambudi"`
	Email    string `json:"email" example:"budipambudi@gmail.com"`
}
