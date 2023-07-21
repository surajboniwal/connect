package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateParam(p any) *AppError {
	err := validate.Struct(p)

	if err == nil {
		return nil
	}

	var appError AppError

	validationErrors := err.(validator.ValidationErrors)

	appError = parseValidationError(&validationErrors[0])

	return &appError
}

func parseValidationError(err *validator.FieldError) AppError {
	var appError AppError = AppError{
		OriginalError: *err,
		Tag:           strings.ToLower((*err).Field()),
		Code:          400,
	}
	switch (*err).Tag() {
	case "required":
		appError.UserMessage = "Required"
	case "email":
		appError.UserMessage = "Pleaase enter a valid email address."
	case "alpha":
		appError.UserMessage = "Invalid input"
	case "alphanum":
		appError.UserMessage = "Invalid input"
	case "e164":
		appError.UserMessage = "Please enter a valid phone number."
	default:
		appError = AppError{
			OriginalError: *err,
			Tag:           "global",
			UserMessage:   "Something went wrong",
		}
	}

	return appError
}
