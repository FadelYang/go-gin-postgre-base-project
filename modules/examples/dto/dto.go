package dto

import "github.com/google/uuid"

type ExampleDTO struct {
	ID   uuid.UUID `json:"id" example:"a53515e3-5a7f-440b-82f6-3d84ac7ce746"`
	Name string    `json:"name" example:"Bakwan Jagung"`
}

type CreateExample struct {
	Name string `json:"name" example:"Bakwan Jagung" binding:"required"`
}
