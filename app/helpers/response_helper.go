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
		} else if tag == "number" {
			message = "Field " + field + " must be a number."
		} else if tag == "gt" {
			message = "Field " + field + " must be a greater than " + param + "."
		}

		errors = append(errors, message)
	}

	return errors
}

func FormatMessageValidationErrors(errors validator.ValidationErrors) []string {
	errorMessages := []string{}

	for _, err := range errors {
		field := err.Field()
		tag := err.Tag()
		param := err.Param()

		message := "Validation error on field " + field + "."

		if tag == "required" {
			message = "Field " + field + " is required."
		} else if tag == "email" {
			message = "Field " + field + " must be a valid email."
		} else if tag == "min" {
			message = "Field " + field + " must be at least " + param + " characters."
		} else if tag == "number" {
			message = "Field " + field + " must be a number."
		} else if tag == "gt" {
			message = "Field " + field + " must be a greater than " + param + "."
		} else if tag == "email_available" {
			message = "Field " + field + " is already in use."
		}

		errorMessages = append(errorMessages, message)
	}

	return errorMessages
}