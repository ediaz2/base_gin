package middleware

import (
	"fmt"
	"net/http"

	"github.com/Oudwins/zog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"tic_tac_boom/pkg/logger"
)

type ValidationSource string

const (
	BodyJSON    ValidationSource = "json"
	QueryParams ValidationSource = "query"
)

type contextKey string

const (
	validatedBodyKey  contextKey = "validated_body"
	validatedQueryKey contextKey = "validated_query"
)

func getContextKey(source ValidationSource) contextKey {
	switch source {
	case BodyJSON:
		return validatedBodyKey
	case QueryParams:
		return validatedQueryKey
	default:
		return ""
	}
}

func Validation[T any](source ValidationSource, schema *zog.StructSchema) gin.HandlerFunc {
	if getContextKey(source) == "" {
		panic(fmt.Sprintf("middleware.Validation: invalid source: %v", source))
	}

	return func(c *gin.Context) {
		var validated T
		var rawData map[string]any

		switch source {
		case BodyJSON:
			if err := c.ShouldBindJSON(&rawData); err != nil {
				logger.Error(c, "JSON binding error", zap.Error(err))
				c.JSON(http.StatusBadRequest, gin.H{
					"error":         "Request binding failed",
					"internal_code": "BINDING_ERROR",
					"details":       err.Error(),
				})
				c.Abort()
				return
			}

		case QueryParams:
			queryParams := c.Request.URL.Query()
			rawData = make(map[string]any)
			for key, values := range queryParams {
				if len(values) == 1 {
					rawData[key] = values[0]
				} else {
					rawData[key] = values
				}
			}

		default:
			panic(fmt.Sprintf("unsupported validation source: %v", source))
		}

		if issues := schema.Parse(rawData, &validated); len(issues) > 0 {
			validationErrors := zog.Issues.SanitizeMap(issues)
			filteredErrors := make(map[string][]string)
			for field, errs := range validationErrors {
				if len(field) > 0 && field[0] != '$' {
					filteredErrors[field] = errs
				}
			}

			logger.Error(c, "validation error", zap.Any("validation_errors", filteredErrors))
			c.JSON(http.StatusBadRequest, gin.H{
				"error":             "Validation failed",
				"internal_code":     "VALIDATION_ERROR",
				"details":           "One or more fields failed validation",
				"validation_errors": filteredErrors,
			})
			c.Abort()
			return
		}

		c.Set(string(getContextKey(source)), &validated)
		c.Next()
	}
}

func GetValidatedBody[T any](c *gin.Context) (*T, error) {
	return getValidatedData[T](c, BodyJSON)
}

func GetValidatedQuery[T any](c *gin.Context) (*T, error) {
	return getValidatedData[T](c, QueryParams)
}

func getValidatedData[T any](c *gin.Context, source ValidationSource) (*T, error) {
	key := getContextKey(source)
	if key == "" {
		return nil, fmt.Errorf("invalid validation source: %v", source)
	}

	if val, exists := c.Get(string(key)); exists {
		if ptr, ok := val.(*T); ok {
			return ptr, nil
		}
		return nil, fmt.Errorf("incorrect type for validated %v data", source)
	}

	return nil, fmt.Errorf("no validated data found for %v - ensure validation middleware is applied", source)
}
