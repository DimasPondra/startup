package structs

import "os"

type userResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url"`
}

func UserResponse(user User, token string) userResponse {
	appUrl := os.Getenv("APP_URL")

	formatter := userResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Role:       user.Role.Name,
		Token:      token,
	}

	if user.FileID != nil {
		formatter.ImageURL = appUrl + "images/" + user.File.Location + "/" + user.File.Name
	}

	return formatter
}
