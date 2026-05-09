package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"project-root/modules/auth/dto"
	"project-root/modules/auth/repository"
	userDTO "project-root/modules/users/dto"
	userModel "project-root/modules/users/model"
	userRepository "project-root/modules/users/repository"
	"project-root/tools"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

var (
	accessSecret  = []byte(os.Getenv("ACCESS_SECRET_KEY"))
	refreshSecret = []byte(os.Getenv("REFRESH_SECRET_KEY"))
)

type AuthUsecase interface {
	Register(form userDTO.CreateUser) (createdUser *userModel.User, err error)
	Login(form dto.LoginDTO) (response *dto.LoginResponse, code int, err error)
}

type authUsecase struct {
	authRepo repository.AuthRepository
	userRepo userRepository.UserRepository
}

func NewAuthUsecase(
	authRepo repository.AuthRepository,
	userRepo userRepository.UserRepository,
) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (u *authUsecase) Register(form userDTO.CreateUser) (createdUser *userModel.User, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}

	createUserForm := userModel.User{
		Username:     form.Username,
		Email:        form.Email,
		Phonenumber:  form.PhoneNumber,
		FirstName:    form.FirstName,
		LastName:     form.LastName,
		PasswordHash: string(hashedPassword),
	}

	user, err := u.userRepo.Create(createUserForm)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				switch pgErr.ConstraintName {

				case "users_username_key":
					return &user, &tools.ValidationError{
						Errors: []tools.FieldError{
							{
								Field:   "username",
								Message: "username already exists",
							},
						},
					}

				case "users_email_key":
					return &user, &tools.ValidationError{
						Errors: []tools.FieldError{
							{
								Field:   "email",
								Message: "email already exists",
							},
						},
					}

				case "users_phonenumber_key":
					return &user, &tools.ValidationError{
						Errors: []tools.FieldError{
							{
								Field:   "phonenumber",
								Message: "phonenumber already exists",
							},
						},
					}
				}
			}
		}

		return &user, err
	}

	return &user, nil
}

func (u *authUsecase) Login(form dto.LoginDTO) (response *dto.LoginResponse, code int, err error) {
	var user userModel.User

	switch form.ChoosenKey {
	case "email":
		user, err = u.userRepo.FindByEmail(form.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	case "username":
		user, err = u.userRepo.FindByUsername(form.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	case "phonenumber":
		user, err = u.userRepo.FindByPhonenumber(form.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	}

	hashedPassword, code, err := u.authRepo.Login(form)
	if err != nil {
		return nil, code, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(form.RawPassword), []byte(hashedPassword)); err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid credentials")
	}

	accessToken, err := u.generateAccessToken(user.ID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	refreshToken, err := u.generateRefreshToken(user.ID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, http.StatusOK, nil
}

func (u *authUsecase) generateAccessToken(userID uuid.UUID) (string, error) {
	jti := uuid.NewString()

	claims := dto.AccessTokenClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (u *authUsecase) generateRefreshToken(userID uuid.UUID) (string, error) {
	jti := uuid.NewString()

	claims := dto.RefreshTokenClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
