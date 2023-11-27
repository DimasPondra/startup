package app

import (
	"startup/app/services"
	"startup/app/structs"

	"github.com/go-playground/validator/v10"
)

func RegisterEmailAvailableValidation(validate *validator.Validate, userService services.UserService) error {
	err := validate.RegisterValidation("email_available", func(fl validator.FieldLevel) bool {
		value, _, _ := fl.ExtractType(fl.Field())

		checkEmail := structs.CheckEmailRequest{
			Email: value.String(),
		}

		emailAvailable, _ := userService.IsEmailAvailable(checkEmail)
		return emailAvailable
	})

	return err
}	