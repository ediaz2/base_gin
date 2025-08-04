package middleware

import (
	"tic_tac_boom/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set("requestID", requestID)
		c.Header(RequestIDKey, requestID)

		loggerWithRequestID := logger.WithRequestID(requestID)
		c.Set(string(logger.LoggerKey), loggerWithRequestID)

		c.Next()
	}
}

func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get("requestID"); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}
