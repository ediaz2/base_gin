package persistence

import (
	"context"
	"tic_tac_boom/internal/core/auth/domain"
	coreErrors "tic_tac_boom/internal/core/errors"
	"tic_tac_boom/internal/infrastructure/supabase"
	"tic_tac_boom/pkg/logger"

	"github.com/supabase-community/gotrue-go/types"
	"go.uber.org/zap"
)

func (r *AuthRepository) LoginUser(ctx context.Context, email, password string) (string, error) {
	resp, err := r.client.Auth.Token(types.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})

	if err != nil {
		logger.Error(ctx, "Failed to login user", zap.Error(err))
		return "", mapSupabaseLoginError(err)
	}

	logger.Info(ctx, "User logged in successfully", zap.String("user_id", resp.User.ID.String()))
	return resp.AccessToken, nil
}

func mapSupabaseLoginError(err error) error {
	supabaseErr := supabase.ParseError(err)
	if supabaseErr == nil {
		return coreErrors.NewInternalError(err)
	}

	switch supabaseErr.Code {
	case 429:
		return coreErrors.NewRateLimitError(err)
	case 400:
		return domain.NewInvalidCredentialsError(err)
	default:
		return coreErrors.NewInternalError(err)
	}
}
