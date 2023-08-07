package middleware

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func AddCacheHeaders(cfg config.Config, c *fiber.Ctx) {
	c.Append("Last-Modified", cfg.LastModified)
	c.Append("Cache-Control", "max-age=2592000")
}
