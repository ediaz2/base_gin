package persistence

import (
	"tic_tac_boom/internal/infrastructure/supabase"
)

type AuthRepository struct {
	client *supabase.SupabaseClient
}

func NewAuthRepository(client *supabase.SupabaseClient) *AuthRepository {
	return &AuthRepository{
		client: client,
	}
}
