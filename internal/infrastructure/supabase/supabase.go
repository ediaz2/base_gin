package supabase

import (
	"encoding/json"
	"strings"
	"tic_tac_boom/internal/infrastructure/config"

	"github.com/supabase-community/supabase-go"
)

type SupabaseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"msg"`
}

type SupabaseClient struct {
	*supabase.Client
}

func NewSupabaseClient(cfg config.SupabaseConfig) (*SupabaseClient, error) {
	client, err := supabase.NewClient(cfg.URL, cfg.APIKey, nil)
	if err != nil {
		return nil, err
	}

	return &SupabaseClient{Client: client}, nil
}

func ParseError(err error) *SupabaseError {
	if err == nil {
		return nil
	}

	raw := err.Error()
	if i := strings.Index(raw, "{"); i != -1 {
		raw = raw[i:]
	}

	var se SupabaseError
	if json.Unmarshal([]byte(raw), &se) == nil {
		return &se
	}
	return nil

}
