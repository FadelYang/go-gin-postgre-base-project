package model

import "github.com/google/uuid"

type Auth struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phonenumber"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
}

type GenerateTokenPayload struct {
	UserID    uuid.UUID `json:"user_id"`
	SessionID uuid.UUID `json:"session_id"`
	Version   uint      `json:"version"`
}
