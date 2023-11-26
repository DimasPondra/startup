package structs

import "os"

type userResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url"`
}

func UserResponse(user User, token string) userResponse {
	appUrl := os.Getenv("APP_URL")

	formatter := userResponse{
		ID: user.ID,
		Name: user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
		Token: token,
	}

	if user.AvatarFileName != "" {
		formatter.ImageURL = appUrl + user.AvatarFileName
	}

	return formatter
}