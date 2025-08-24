package service

import (
	"fmt"
	"project-root/modules/examples/dto"
	"project-root/modules/examples/model"
	"project-root/modules/examples/repository"
)

type ExampleService interface {
	GetExamples() ([]dto.ExampleDTO, error)
	CreateExample(example dto.ExampleDTO) (dto.ExampleDTO, error)
}

type exampleService struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleService(example repository.ExampleRepository) ExampleService {
	return &exampleService{
		exampleRepository: example,
	}
}

func (s *exampleService) GetExamples() ([]dto.ExampleDTO, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *exampleService) CreateExample(example dto.ExampleDTO) (dto.ExampleDTO, error) {
	exampleForm := model.Example{
		Name: example.Name,
	}

	createdExample, err := s.exampleRepository.Create(exampleForm)
	if err != nil {
		return dto.ExampleDTO{}, err
	}

	return dto.ExampleDTO{
		Name: createdExample.Name,
	}, fmt.Errorf("not implemented")
}
