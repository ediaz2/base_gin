package persistence

import (
	"context"
	"tic_tac_boom/internal/core/auth/domain"
	"tic_tac_boom/internal/core/errors"
	"tic_tac_boom/internal/infrastructure/supabase"
	"tic_tac_boom/pkg/logger"

	"github.com/supabase-community/gotrue-go/types"
	"go.uber.org/zap"
)

func (r *AuthRepository) RegisterUser(ctx context.Context, user *domain.User) (string, error) {
	resp, err := r.client.Auth.Signup(types.SignupRequest{
		Email:    user.Email,
		Password: user.Password,
		Data: map[string]any{
			"username":     user.Username,
			"display_name": user.DisplayName,
		},
	})

	if err != nil {
		return "", mapSupabaseError(err)
	}

	userID := resp.ID.String()
	logger.Info(ctx, "User created successfully in Supabase", zap.String("user_id", userID))

	return userID, nil
}

func mapSupabaseError(err error) error {
	supabaseErr := supabase.ParseError(err)

	if supabaseErr == nil {
		return errors.NewInternalError(err)
	}

	switch supabaseErr.Code {
	case 429:
		return errors.NewRateLimitError(err)
	default:
		return errors.NewInternalError(err)
	}
}
