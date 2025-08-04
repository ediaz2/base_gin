package server

import (
	"context"
	"fmt"
	"tic_tac_boom/internal/infrastructure/config"
	"tic_tac_boom/internal/infrastructure/middleware"
	"tic_tac_boom/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HTTPServer struct {
	app    *gin.Engine
	config config.ServerConfig
}

func NewHTTPServer(cfg config.ServerConfig, routeGroups ...RouteGroup) *HTTPServer {
	app := gin.New()

	if err := app.SetTrustedProxies([]string{"192.168.1.0/24"}); err != nil {
		ctx := context.Background()
		logger.Fatal(ctx, "Failed to set trusted proxies", zap.Error(err))
	}

	app.Use(middleware.RequestID())
	// app.Use(gin.Logger())
	app.Use(middleware.Helmet(middleware.HelmetConfig{}))
	app.Use(middleware.ErrorHandler())

	for _, routeGroup := range routeGroups {
		routeGroup.RegisterRoutes(app)
	}

	return &HTTPServer{
		app:    app,
		config: cfg,
	}
}

func (s *HTTPServer) Start() error {
	return s.app.Run(fmt.Sprintf(":%s", s.config.Port))
}
