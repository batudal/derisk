package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

func HandleLazyImage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("components/common/lazy-image", fiber.Map{
			"ImageURL": c.Params("imageURL"),
		})
	}
}
