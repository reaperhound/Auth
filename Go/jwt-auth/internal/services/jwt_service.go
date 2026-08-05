package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServie struct{}

func NewJwtService() *JwtServie {
	return &JwtServie{}
}

const (
	accessSecret  = "ACCESS_SECRET"
	refreshSecret = "REFRESH_SECRET"
)

func (s *JwtServie) GenerateAccessTok(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
		"type":     "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

func (s *JwtServie) GenerateRefreshTok(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":      time.Now(),
		"type":     "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshSecret))
}
