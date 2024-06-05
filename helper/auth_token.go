package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthTokenService interface {
	GenerateToken(payload interface{}) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	var JWT_SECRET = os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
