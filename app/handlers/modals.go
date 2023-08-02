package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func BetaList(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("components/modals/beta-list", fiber.Map{})
	}
}
