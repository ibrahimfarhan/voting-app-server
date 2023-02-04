package models

import "net/http"

const (
	ResourceDoesNotExist = "Resource does not exist"
	DuplicateData        = "Duplicate Data"
	SomethingWentWrong   = "Something went wrong"
	ValidationError      = "Validation Error"
	Unauthorized         = "Unauthorized"
	Forbidden            = "Forbidden"
	ResourceUnmodified   = "Resource was not modified"
	ExpiredToken         = "Expired Token"
	AuthRequired         = "Authentication is required"
)

type ServerError struct {
	Message    string `json:"message,omitempty"`
	StatusText string `json:"statusText,omitempty"`
}

func NewServerError(message string, status int) *ServerError {
	return &ServerError{Message: message, StatusText: http.StatusText(status)}
}
