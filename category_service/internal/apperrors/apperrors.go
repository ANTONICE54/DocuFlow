package apperrors

import "net/http"

type Type string

const (
	BadRequest   Type = "BAD_REQUEST"
	Conflict     Type = "CONFLICT"
	Internal     Type = "INTERNAL"
	NotFound     Type = "NOT_FOUND"
	Unauthorized Type = "Unauthorized"
)

type AppError struct {
	JSONResponse map[string]interface{}
	Message      string
	Type         Type
}

func NewError(msg string, code Type, jsonResp string) *AppError {
	return &AppError{
		Message:      msg,
		Type:         code,
		JSONResponse: map[string]interface{}{"error": jsonResp},
	}
}

func (err *AppError) Error() string {
	return err.Message
}

func (err *AppError) Status() int {
	switch err.Type {
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case Unauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func ErrDatabase(msg string) *AppError {
	return NewError(msg, Internal, "Database error raised an error")
}

func ErrInternalServer(msg string) *AppError {
	return NewError(msg, Internal, "Internal server error")
}

func ErrBadRequest(msg string) *AppError {
	return NewError(msg, Internal, "Bad request")
}

func ErrUnauthorized(msg string) *AppError {
	return NewError(msg, Unauthorized, "Unauthorized")
}

func ErrNotFound(msg string) *AppError {
	return NewError(msg, NotFound, "Category not found")
}
