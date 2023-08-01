package routes

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func Listen(cfg config.Config) {
	cfg.App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	cfg.App.Listen(":3000")
}
