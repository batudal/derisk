package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func Index(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Modal":        c.Query("modal"),
			"Email":        c.Query("email"),
			"CustomerType": c.Query("customer_type"),
		}, "layouts/public")
	}
}

func Blog(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("pages/blog", fiber.Map{}, "layouts/public")
	}
}
