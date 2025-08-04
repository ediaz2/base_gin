package repositories

import (
	"context"
	"tic_tac_boom/internal/core/auth/domain"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, user *domain.User) (string, error)
	LoginUser(ctx context.Context, email, password string) (string, error)
}
