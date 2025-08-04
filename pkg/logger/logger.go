package logger

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	LoggerKey    contextKey = "logger"
)

var baseLogger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	baseLogger, _ = config.Build(zap.AddCallerSkip(1))
}

func GetBaseLogger() *zap.Logger {
	return baseLogger
}

func WithRequestID(requestID string) *zap.Logger {
	return baseLogger.With(zap.String("request_id", requestID))
}

func FromContext(ctx context.Context) *zap.Logger {
	switch c := ctx.(type) {
	case *gin.Context:
		if logger, exists := c.Get(string(LoggerKey)); exists {
			if zapLogger, ok := logger.(*zap.Logger); ok {
				return zapLogger
			}
		}
	case context.Context:
		if logger := c.Value(LoggerKey); logger != nil {
			if zapLogger, ok := logger.(*zap.Logger); ok {
				return zapLogger
			}
		}
	}
	return baseLogger
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Info(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Error(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Warn(msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Debug(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	FromContext(ctx).Fatal(msg, fields...)
}

func Sync() error {
	return baseLogger.Sync()
}
