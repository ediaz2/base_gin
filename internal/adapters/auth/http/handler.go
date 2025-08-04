package http

import (
	"tic_tac_boom/internal/core/auth/application/commands"
)

type AuthHandler struct {
	registerUserCmd *commands.RegisterUserCommand
}

func NewAuthHandler(registerUserCmd *commands.RegisterUserCommand) *AuthHandler {
	return &AuthHandler{
		registerUserCmd: registerUserCmd,
	}
}
