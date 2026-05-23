package repository

import (
	"context"
	"project-root/modules/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByUsername(ctx context.Context, username string) (model.User, error)
	FindByPhonenumber(ctx context.Context, phonenumber string) (model.User, error)
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

func (r *userRepository) Create(ctx context.Context, User model.User) (model.User, error) {
	if err := r.db.
		WithContext(ctx).
		Create(&User).Error; err != nil {
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

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByPhonenumber(ctx context.Context, phonenumber string) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "phonenumber = ?", phonenumber).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
