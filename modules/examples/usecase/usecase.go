package usecase

import (
	"fmt"
	"project-root/modules/examples/dto"
	"project-root/modules/examples/model"
	"project-root/modules/examples/repository"
)

type ExampleUsecase interface {
	GetExamples() ([]dto.ExampleDTO, error)
	CreateExample(example dto.ExampleDTO) (dto.ExampleDTO, error)
}

type exampleUsecase struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleUsecase(example repository.ExampleRepository) ExampleUsecase {
	return &exampleUsecase{
		exampleRepository: example,
	}
}

func (s *exampleUsecase) GetExamples() ([]dto.ExampleDTO, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *exampleUsecase) CreateExample(example dto.ExampleDTO) (dto.ExampleDTO, error) {
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
