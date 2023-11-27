package structs

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Name       string `json:"name" validate:"required"`
	Occupation string `json:"occupation" validate:"required"`
	Email      string `json:"email" validate:"required,email,email_available"`
	Password   string `json:"password" validate:"required,min=6"`
}

type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
