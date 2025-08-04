package commands

import (
	"context"
	"tic_tac_boom/internal/core/auth/domain"
	"tic_tac_boom/internal/core/auth/ports/dto"
	"tic_tac_boom/internal/core/auth/ports/repositories"
	"tic_tac_boom/pkg/logger"

	"go.uber.org/zap"
)

type RegisterUserCommand struct {
	authRepo repositories.AuthRepository
}

func NewRegisterUserCommand(authRepo repositories.AuthRepository) *RegisterUserCommand {
	return &RegisterUserCommand{
		authRepo: authRepo,
	}
}

func (cmd *RegisterUserCommand) Execute(ctx context.Context, req dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	logger.Info(ctx, "Executing register user command",
		zap.String("email", req.Email),
		zap.String("username", req.Username),
	)

	user, err := domain.NewUser(req.Email, req.Username, req.Password)
	if err != nil {
		logger.Error(ctx, "Failed to create user domain object", zap.Error(err))
		return nil, err
	}

	logger.Debug(ctx, "User domain object created successfully")

	userID, err := cmd.authRepo.RegisterUser(ctx, user)
	if err != nil {
		logger.Error(ctx, "Failed to register user in repository", zap.Error(err))
		return nil, err
	}

	logger.Info(ctx, "User registered successfully",
		zap.String("user_id", userID),
	)

	return &dto.RegisterUserResponse{
		UserID:  userID,
		Message: "User registered successfully",
	}, nil
}
