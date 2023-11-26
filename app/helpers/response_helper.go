package helpers

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseAPI(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		tag := e.ActualTag() // tag ex: required, min
		field := e.Field() // field ex: Occupation, Password
		param := e.Param() // field ex: 6(this value from min)

		message := "Validation error on field " + field

		if tag == "required" {
			message = "Field " + field + " is required."
		} else if tag == "email" {
			message = "Field " + field + " must be a valid email."
		} else if tag == "min" {
			message = "Field " + field + " must be at least " + param + " characters."
		}

		errors = append(errors, message)
	}

	return errors
}