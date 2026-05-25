package services

import (
	"errors"
	"project-root/modules/auth/dto"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	accessSecret  string
	refreshSecret string
}

func (s *JWTService) ValidateAccessToken(
	tokenString string,
) (*dto.AccessTokenClaim, error) {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return []byte(s.accessSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*dto.AccessTokenClaim)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
