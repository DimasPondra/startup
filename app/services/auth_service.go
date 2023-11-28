package services

import (
	"errors"
	"os"
	"startup/app/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(user structs.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct{}

func NewAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(user structs.User) (string, error) {
	claim := jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role.Name,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secretKey := os.Getenv("JWT_SECRET_KEY")

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
