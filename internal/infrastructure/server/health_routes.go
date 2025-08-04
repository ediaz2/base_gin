package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthRoutes struct {
	version   string
	startTime time.Time
}

func NewHealthRoutes(version string) *HealthRoutes {
	return &HealthRoutes{
		version:   version,
		startTime: time.Now(),
	}
}

func (h *HealthRoutes) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", h.healthCheck)
}

func (h *HealthRoutes) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"message":   "Service is running",
		"version":   h.version,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"uptime":    time.Since(h.startTime).String(),
	})
}
