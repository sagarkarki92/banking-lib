package err

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e *AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func NewValidationError(message string) *AppError {
	return &AppError{Code: http.StatusUnprocessableEntity, Message: message}
}

func NewUnauthorizedError() *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: "Unauthorized!"}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}
