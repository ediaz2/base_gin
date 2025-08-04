package http

import (
	"net/http"
	"tic_tac_boom/internal/core/auth/ports/dto"
	"tic_tac_boom/internal/infrastructure/middleware"
	"tic_tac_boom/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	logger.Info(c, "Starting user registration request")

	req, err := middleware.GetValidatedBody[dto.RegisterUserRequest](c)

	if err != nil {
		logger.Error(c, "Failed to get validated request body", zap.Error(err))
		_ = c.Error(err)
		return
	}

	response, err := h.registerUserCmd.Execute(c, *req)
	if err != nil {
		logger.Error(c, "Command execution failed", zap.Error(err))
		_ = c.Error(err)
		return
	}

	logger.Info(c, "User registration completed successfully")
	c.JSON(http.StatusCreated, response)
}
