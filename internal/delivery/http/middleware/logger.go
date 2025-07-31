package middleware

import (
	"dubai-auto/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func ZerologMiddleware(logger *logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if c.Response().StatusCode() != 200 {
			logger.Info().
				Str("method", c.Method()).
				Str("path", c.Path()).
				Int("status", c.Response().StatusCode()).
				Msg("Non-200 response")
		}
		return err
	}
}
