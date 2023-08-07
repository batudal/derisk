package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func HandleIndexPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Append("Cache-Control", "no-cache, no-store, must-revalidate")
		return c.Render("pages/index", fiber.Map{
			"Modal":        c.Query("modal"),
			"Email":        c.Query("email"),
			"CustomerType": c.Query("customer_type"),
		}, "layouts/public")
	}
}

func HandleAboutPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("pages/about", fiber.Map{}, "layouts/public")
	}
}
