package dto

import "github.com/google/uuid"

type ExampleDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
