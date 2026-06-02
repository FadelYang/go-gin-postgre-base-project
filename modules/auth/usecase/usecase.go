package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"project-root/internal/services"
	"project-root/modules/auth/dto"
	"project-root/modules/auth/model"
	"project-root/modules/auth/repository"
	userDTO "project-root/modules/users/dto"
	userModel "project-root/modules/users/model"
	userRepository "project-root/modules/users/repository"
	"project-root/tools"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, form userDTO.CreateUser) (createdUser *userModel.User, err error)
	Login(ctx context.Context, form dto.LoginDTO) (response *dto.LoginResponse, code int, err error)
	RefreshLogin(ctx context.Context, refteshToken string) (newAccessToken *string, err error)
	Logout(ctx context.Context, form dto.Logout) (code int, err error)
}

type authUsecase struct {
	RedisClient redis.Client
	authRepo    repository.AuthRepository
	userRepo    userRepository.UserRepository
	jwtService  *services.JWTService
}

func NewAuthUsecase(
	redisClient *redis.Client,
	authRepo repository.AuthRepository,
	userRepo userRepository.UserRepository,
	jwtService *services.JWTService,
) AuthUsecase {
	return &authUsecase{
		RedisClient: *redisClient,
		authRepo:    authRepo,
		userRepo:    userRepo,
		jwtService:  jwtService,
	}
}

func (u *authUsecase) Register(ctx context.Context, form userDTO.CreateUser) (createdUser *userModel.User, err error) {
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

	user, err := u.userRepo.Create(ctx, createUserForm)
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

func (u *authUsecase) Login(ctx context.Context, form dto.LoginDTO) (response *dto.LoginResponse, code int, err error) {
	var user userModel.User

	switch form.ChoosenKey {
	case "email":
		user, err = u.userRepo.FindByEmail(ctx, form.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	case "username":
		user, err = u.userRepo.FindByUsername(ctx, form.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	case "phonenumber":
		user, err = u.userRepo.FindByPhonenumber(ctx, form.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	}

	hashedPassword, code, err := u.authRepo.Login(ctx, form)
	if err != nil {
		return nil, code, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(form.RawPassword)); err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid credentials")
	}

	sessionID := uuid.New()

	generateTokenPayload := model.GenerateTokenPayload{
		UserID:    user.ID,
		SessionID: sessionID,
		Version:   uint(user.TokenVersion),
		Role:      user.Role.Name,
	}

	accessToken, err := u.jwtService.GenerateAccessToken(generateTokenPayload)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	refreshToken, err := u.jwtService.GenerateRefreshToken(generateTokenPayload)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	hashedRefreshToken := tools.HashToken(refreshToken)

	redisKey := fmt.Sprintf(
		"refresh_token:%s:%s",
		user.ID,
		sessionID,
	)

	err = u.RedisClient.Set(
		ctx,
		redisKey,
		hashedRefreshToken,
		7*24*time.Hour,
	).Err()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, http.StatusOK, nil
}

func (u *authUsecase) RefreshLogin(ctx context.Context, refreshToken string) (newAccessToken *string, err error) {
	token, err := u.jwtService.ParseRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*dto.RefreshTokenClaim)

	if claims.Type != "refresh" {
		return nil, errors.New("invalid token type")
	}

	userDetail, err := u.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	if userDetail.TokenVersion != int(claims.Version) {
		return nil, errors.New("token already revoked")
	}

	redisKey := fmt.Sprintf(
		"refresh_token:%s:%s",
		claims.UserID,
		claims.SessionID,
	)

	storedHash, err := u.RedisClient.Get(ctx, redisKey).Result()
	if err != nil {
		return nil, errors.New("refresh token already revoked, please login again")
	}

	incomingHash := tools.HashToken(refreshToken)

	if incomingHash != storedHash {
		return nil, errors.New("invalid refresh token")
	}

	accessTokenPayload := model.GenerateTokenPayload{
		UserID:    claims.UserID,
		SessionID: claims.SessionID,
		Version:   claims.Version,
	}

	accessToken, err := u.jwtService.GenerateAccessToken(accessTokenPayload)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func (u *authUsecase) Logout(ctx context.Context, form dto.Logout) (code int, err error) {
	token, err := u.jwtService.ParseRefreshToken(ctx, form.RefreshToken)
	if err != nil {
		return http.StatusBadRequest, err
	}

	claims := token.Claims.(*dto.RefreshTokenClaim)

	redisKey := fmt.Sprintf(
		"refresh_token:%s:%s",
		claims.UserID,
		claims.SessionID,
	)

	u.RedisClient.Del(ctx, redisKey)

	updateTokenVersionPayload := dto.LoginDTO{
		Key:        claims.UserID.String(),
		ChoosenKey: "userid",
	}

	_, code, err = u.authRepo.UpdateTokenVersion(ctx, updateTokenVersionPayload)
	if err != nil {
		return code, err
	}

	return code, nil
}
