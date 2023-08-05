package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func HandleBetaList(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("components/modals/beta-list", fiber.Map{})
	}
}

func HandleFeedback(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("components/modals/feedback", fiber.Map{
			"Email":        c.Query("email"),
			"CustomerType": c.Query("customer_type"),
		})
	}
}

func HandleMobileMenu(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("components/modals/mobile-menu", fiber.Map{})
	}
}
