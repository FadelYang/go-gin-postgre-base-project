package model

import "github.com/google/uuid"

type Auth struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phonenumber"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
}
