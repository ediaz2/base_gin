package server

import (
	"tic_tac_boom/internal/adapters/auth/http"
	"tic_tac_boom/internal/adapters/auth/http/validation"
	"tic_tac_boom/internal/core/auth/ports/dto"
	"tic_tac_boom/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authHandler *http.AuthHandler
}

func NewAuthRoutes(authHandler *http.AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		authHandler: authHandler,
	}
}

func (a *AuthRoutes) RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	{
		routes.POST("/register",
			middleware.Validation[dto.RegisterUserRequest](middleware.BodyJSON, validation.RegisterUserRequestSchema),
			a.authHandler.RegisterUser,
		)
	}
}
