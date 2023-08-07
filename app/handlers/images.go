package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func HandleHeroImage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("components/landing/hero-image", fiber.Map{})
	}
}
