package services

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct{}

func NewAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secretKey := os.Getenv("JWT_SECRET_KEY")

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}