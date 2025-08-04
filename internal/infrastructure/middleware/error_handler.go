package middleware

import (
	"errors"
	"net/http"
	coreErrors "tic_tac_boom/internal/core/errors"
	"tic_tac_boom/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			var domainErr *coreErrors.DomainError
			if errors.As(err.Err, &domainErr) {
				c.JSON(domainErr.Code, gin.H{
					"error":         domainErr.Message,
					"internal_code": domainErr.InternalCode,
				})
				if domainErr.Cause != nil {
					logger.Error(c, "Domain error cause", zap.Error(domainErr.Cause))
				} else {
					logger.Error(c, "Domain error without cause", zap.Error(domainErr))
				}
				c.Abort()
				return
			}

			logger.Error(c, "Internal server error", zap.Error(err.Err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			c.Abort()
		}
	}
}
