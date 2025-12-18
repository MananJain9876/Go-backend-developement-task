package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-Id"

// RequestID injects a requestId header into responses (and sets it on the request if missing).
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqID := c.Get(requestIDHeader)
		if reqID == "" {
			reqID = uuid.New().String()
		}
		c.Set(requestIDHeader, reqID)
		return c.Next()
	}
}

// RequestLogger logs request method, path, status, and duration.
func RequestLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		logger.Info("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("request_id", c.Get(requestIDHeader)),
		)

		return err
	}
}


