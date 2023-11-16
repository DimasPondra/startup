package structs

import (
	"os"
	"time"
)

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RegisterRequest struct {
	Name 		string	`json:"name" binding:"required"`
	Occupation  string	`json:"occupation" binding:"required"`
	Email 		string	`json:"email" binding:"required,email"`
	Password 	string	`json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required,min=6"`
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	Occupation 	string 	`json:"occupation"`
	Email 		string 	`json:"email"`
	Token 		string 	`json:"token"`
	ImageURL	string	`json:"image_url"`
}

func UserResponse(user User, token string) userResponse {
	appUrl := os.Getenv("APP_URL")

	formatter := userResponse{
		ID: user.ID,
		Name: user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
		Token: token,
		ImageURL: appUrl + user.AvatarFileName,
	}

	return formatter
}