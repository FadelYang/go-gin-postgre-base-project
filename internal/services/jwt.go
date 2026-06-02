package services

import (
	"context"
	"errors"
	"fmt"
	"project-root/modules/auth/dto"
	"project-root/modules/auth/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	accessSecret  string
	refreshSecret string
}

func NewJWTService(accessSecret string, refreshSecret string) *JWTService {
	return &JWTService{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

func (s *JWTService) ValidateAccessToken(
	tokenString string,
) (*dto.AccessTokenClaim, error) {

	claims := &dto.AccessTokenClaim{}

	fmt.Println("accessSecret: ", []byte(s.accessSecret))

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return []byte(s.accessSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *JWTService) GenerateAccessToken(payload model.GenerateTokenPayload) (string, error) {
	jti := uuid.NewString()

	claims := dto.AccessTokenClaim{
		UserID: payload.UserID,
		Role:   payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * 15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.accessSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *JWTService) GenerateRefreshToken(payload model.GenerateTokenPayload) (string, error) {
	jti := uuid.NewString()

	claims := dto.RefreshTokenClaim{
		UserID:    payload.UserID,
		Role:      payload.Role,
		SessionID: payload.SessionID,
		Version:   payload.Version,
		Type:      "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.refreshSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *JWTService) ParseRefreshToken(ctx context.Context, refreshToken string) (parsedToken *jwt.Token, err error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&dto.RefreshTokenClaim{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(s.refreshSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	return token, nil
}
