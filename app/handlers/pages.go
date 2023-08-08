package handlers

import (
	"github.com/batudal/derisk/app/config"
	"github.com/gofiber/fiber/v2"
)

type Pricing struct {
	CustomerType string
	TagColor     string
	Price        string
	Highlights   []string
	Features     []string
	CTA          string
}

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

func HandleBlogPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("pages/about", fiber.Map{}, "layouts/public")
	}
}

func HandlePricingPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tiers := []Pricing{
			{
				CustomerType: "Solopreneurs",
				TagColor:     "green",
				Price:        "$0/mo",
				Highlights: []string{
					"Unlimited number of users",
					"Unlimited number of projects",
				},
				Features: []string{
					"Up to 5 projects",
					"Up to 5 users",
				},
				CTA: "Join beta list",
			},
			{
				CustomerType: "Startups",
				TagColor:     "cyan",
				Price:        "$99/mo",
				Highlights: []string{
					"Unlimited number of users",
					"Unlimited number of projects",
				},
				Features: []string{
					"Up to 5 projects",
					"Up to 5 users",
				},
				CTA: "Join beta list",
			},
			{
				CustomerType: "Incubators",
				TagColor:     "yellow",
				Price:        "Custom",
				Highlights: []string{
					"Unlimited number of users",
					"Unlimited number of projects",
				},
				Features: []string{
					"Up to 5 projects",
					"Up to 5 users",
				},
				CTA: "Contact us",
			},
		}
		return c.Render("pages/pricing", fiber.Map{
			"Tiers": tiers,
		}, "layouts/public")
	}
}

func HandleContactPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("pages/contact", fiber.Map{}, "layouts/public")
	}
}
