package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type LoginDTO struct {
	Key            string `json:"key"`
	ChoosenKey     string `json:"choosen_key"`
	RawPassword    string `json:"password"`
	HashedPassword string `json:"-"`
}

type RegisterDTO struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phonenumber"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	RawPassword     string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	HashedPassword  string `json:"-"`
}

type AccessTokenClaim struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshTokenClaim struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}
