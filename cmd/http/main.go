package main

import (
	"context"
	"tic_tac_boom/internal/adapters/auth/http"
	"tic_tac_boom/internal/adapters/auth/persistence"
	"tic_tac_boom/internal/core/auth/application/commands"
	"tic_tac_boom/internal/infrastructure/config"
	"tic_tac_boom/internal/infrastructure/server"
	"tic_tac_boom/internal/infrastructure/supabase"
	"tic_tac_boom/pkg/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	defer func() {
		_ = logger.Sync()
	}()

	if err := godotenv.Load(); err != nil {
		logger.Warn(ctx, "Warning: .env file not found", zap.Error(err))
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(ctx, "Failed to load configuration", zap.Error(err))
	}

	supabaseClient, err := supabase.NewSupabaseClient(cfg.Supabase)
	if err != nil {
		logger.Fatal(ctx, "Failed to initialize Supabase client", zap.Error(err))
	}

	authRepo := persistence.NewAuthRepository(supabaseClient)
	registerUserCmd := commands.NewRegisterUserCommand(authRepo)
	authHandler := http.NewAuthHandler(registerUserCmd)

	healthRoutes := server.NewHealthRoutes("1.0.0")
	authRoutes := server.NewAuthRoutes(authHandler)

	httpServer := server.NewHTTPServer(cfg.Server, healthRoutes, authRoutes)

	logger.Info(ctx, "Starting server", zap.String("port", cfg.Server.Port))
	if err := httpServer.Start(); err != nil {
		logger.Fatal(ctx, "Failed to start server", zap.Error(err))
	}
}
