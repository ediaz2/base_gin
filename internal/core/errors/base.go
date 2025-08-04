package errors

import (
	"fmt"
	"net/http"
)

type DomainError struct {
	Code         int
	InternalCode string
	Message      string
	Cause        error
}

func (e *DomainError) Error() string {
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

func NewInternalError(cause error) *DomainError {
	return &DomainError{
		Code:         http.StatusInternalServerError,
		InternalCode: "ERR_INTERNAL_001",
		Message:      "Internal server error occurred",
		Cause:        cause,
	}
}

func NewRateLimitError(cause error) *DomainError {
	return &DomainError{
		Code:         http.StatusTooManyRequests,
		InternalCode: "ERR_RATE_LIMIT_001",
		Message:      "Rate limit exceeded. Please wait before trying again",
		Cause:        cause,
	}
}

func NewNotFoundError(resource string) *DomainError {
	return &DomainError{
		Code:         http.StatusNotFound,
		InternalCode: "ERR_NOT_FOUND_001",
		Message:      fmt.Sprintf("%s not found", resource),
	}
}

func NewAlreadyExistsError(resource string) *DomainError {
	return &DomainError{
		Code:         http.StatusConflict,
		InternalCode: "ERR_ALREADY_EXISTS_001",
		Message:      fmt.Sprintf("%s already exists", resource),
	}
}

func NewInvalidDataError(message string) *DomainError {
	return &DomainError{
		Code:         http.StatusBadRequest,
		InternalCode: "ERR_INVALID_DATA_001",
		Message:      fmt.Sprintf("Invalid data: %s", message),
	}
}

func NewUnauthorizedError(message string) *DomainError {
	return &DomainError{
		Code:         http.StatusUnauthorized,
		InternalCode: "ERR_UNAUTHORIZED_001",
		Message:      fmt.Sprintf("Unauthorized: %s", message),
	}
}

func NewForbiddenError(message string) *DomainError {
	return &DomainError{
		Code:         http.StatusForbidden,
		InternalCode: "ERR_FORBIDDEN_001",
		Message:      fmt.Sprintf("Forbidden: %s", message),
	}
}
