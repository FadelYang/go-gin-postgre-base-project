package repository

import (
	"errors"
	"net/http"
	"project-root/modules/auth/dto"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(form dto.RegisterDTO) (code int, err error)
	Login(form dto.LoginDTO) (passwordHash string, code int, err error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(form dto.RegisterDTO) (code int, err error) {
	baseQuery := `
		INSERT INTO users (username, email, phonenumber, first_name, last_name, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	if err := r.db.Exec(
		baseQuery,
		form.Username,
		form.Email,
		form.PhoneNumber,
		form.FirstName,
		form.LastName,
		form.HashedPassword,
	); err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusCreated, nil
}

func (r *authRepository) Login(form dto.LoginDTO) (passwordHash string, code int, err error) {
	var checkIsExistsQuery string

	switch form.ChoosenKey {
	case "email":
		checkIsExistsQuery = qGetPasswordHashByEmail
	case "phonenumber":
		checkIsExistsQuery = qGetPasswordHashByPhoneNumber
	case "username":
		checkIsExistsQuery = qGetPasswordHashByUsername
	default:
		return "", http.StatusBadRequest, errors.New("unknown login method")
	}

	if err := r.db.Raw(
		checkIsExistsQuery,
		form.Key,
	).Scan(&passwordHash).Error; err != nil {
		return "", http.StatusBadRequest, err
	}

	if passwordHash == "" {
		return "", http.StatusNotFound, errors.New("user not found")
	}

	return passwordHash, http.StatusOK, nil
}
