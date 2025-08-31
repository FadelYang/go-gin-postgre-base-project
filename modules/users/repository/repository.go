package repository

import (
	"project-root/modules/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewuserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Create(User model.User) (model.User, error) {
	if err := r.db.Create(&User).Error; err != nil {
		return model.User{}, err
	}

	return User, nil
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	if err := r.db.Model(&user).Updates(map[string]any{
		"username": user.Username,
		"email":    user.Email,
	}).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByID(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
