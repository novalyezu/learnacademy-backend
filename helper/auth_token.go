package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthTokenService interface {
	GenerateToken(payload interface{}) (string, error)
}

type jwtService struct {
}

func NewAuthTokenService() AuthTokenService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(payload interface{}) (string, error) {
	var JWT_SECRET = os.Getenv("JWT_SECRET")

	claim := jwt.MapClaims{
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
