package domain

import (
	"net/http"
	"tic_tac_boom/internal/core/errors"
)

func NewInvalidCredentialsError(cause error) *errors.DomainError {
	return &errors.DomainError{
		Code:    http.StatusUnauthorized,
		Message: "Invalid credentials provided",
		Cause:   cause,
	}
}
