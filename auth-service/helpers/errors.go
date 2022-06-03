package helpers

import "net/http"

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"statusCode,omitempty"`
}

func (c AppError) AsMessage() *AppError {
	return &AppError{
		Message: c.Message,
		Code:    c.Code,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}
