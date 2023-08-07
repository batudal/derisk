package middleware

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func AddCacheHeaders(cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Append("Last-Modified", cfg.LastModified)
		return c.Next()
	}
}
