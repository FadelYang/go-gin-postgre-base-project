package service

import (
	"project-root/modules/users/dto"
	"project-root/modules/users/model"
	"project-root/modules/users/repository"
	"project-root/tools"

	"github.com/google/uuid"
)

type UserService interface {
	GetAll() ([]dto.UserDTO, error)
	Create(user dto.CreateUser) (dto.UserDTO, error)
	Update(user dto.UpdateUser, userID uuid.UUID) (dto.UserDTO, error)
	Delete(id uuid.UUID) (dto.UserDTO, error)
	FindByID(id uuid.UUID) (dto.UserDTO, error)
	FindByEmail(email string) (dto.UserDTO, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAll() ([]dto.UserDTO, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return []dto.UserDTO{}, err
	}

	result := make([]dto.UserDTO, 0, len(users))
	for _, i := range users {
		result = append(result, dto.UserDTO{
			Username:  i.Username,
			Email:     i.Email,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		})
	}

	return result, nil
}

func (s *userService) Create(form dto.CreateUser) (dto.UserDTO, error) {
	passwordHash, err := tools.HashPassword(form.Password)
	if err != nil {
		return dto.UserDTO{}, err
	}

	userForm := model.User{
		Username:     form.Username,
		Email:        form.Email,
		PasswordHash: passwordHash,
	}

	createdUser, err := s.userRepo.Create(userForm)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserDTO{
		UUID:      createdUser.ID,
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *userService) Update(user dto.UpdateUser, userID uuid.UUID) (dto.UserDTO, error) {
	existData, err := s.userRepo.FindByID(userID)
	if err != nil {
		return dto.UserDTO{}, err
	}

	existData.Username = user.Username
	existData.Email = user.Email

	updatedData, err := s.userRepo.Update(existData)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserDTO{
		UUID:      updatedData.ID,
		Username:  updatedData.Username,
		Email:     updatedData.Email,
		CreatedAt: updatedData.CreatedAt,
		UpdatedAt: updatedData.UpdatedAt,
	}, nil
}

func (s *userService) Delete(id uuid.UUID) (dto.UserDTO, error) {
	userToDelete, err := s.FindByID(id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	if err := s.userRepo.Delete(userToDelete.UUID); err != nil {
		return dto.UserDTO{}, nil
	}

	return userToDelete, nil
}

func (s *userService) FindByID(id uuid.UUID) (dto.UserDTO, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserDTO{
		UUID:      user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) FindByEmail(email string) (dto.UserDTO, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserDTO{
		UUID:      user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
