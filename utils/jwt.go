package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	Generate(userID string, role string) (string, error)
}

type jwtService struct {
	secret string
}

func NewJWT(secret string) JWTService {
	return &jwtService{secret: secret}
}

func (j *jwtService) Generate(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

var ErrInvalidCredentials = errors.New("invalid email or password")
