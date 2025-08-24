package repository

import (
	"fmt"
	"project-root/modules/examples/model"

	"gorm.io/gorm"
)

type ExampleRepository interface {
	FindAll() ([]model.Example, error)
	Create(example model.Example) (model.Example, error)
}

type exampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &exampleRepository{
		db: db,
	}
}

func (r *exampleRepository) FindAll() ([]model.Example, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *exampleRepository) Create(example model.Example) (model.Example, error) {
	// In repository pattern, it is often use model rather than DTO
	var createdExample model.Example

	return createdExample, fmt.Errorf("not implemented")
}
