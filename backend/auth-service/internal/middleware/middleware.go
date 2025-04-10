package middleware

//TODO: Not using this one yet
import (
	"context"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type contextKey string

const httpRequestKey contextKey = "httpRequest"

func LoggerMiddleware(log logger.LoggerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), httpRequestKey, c.Request)
		c.Request = c.Request.WithContext(ctx)

		// Optionally log the request here
		log.Info(ctx, "Incoming request", c.FullPath(), map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})

		c.Next()
	}
}
