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
	Features     []Feature
	CTA          string
}

type Feature struct {
	Title   string
	Enabled bool
	Starred bool
}

func HandleIndexPage(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
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
					"Create up to 5 projects",
					"Create up to 5 tests per project",
				},
				Features: []Feature{
					{"Analytics dashboard", true, false},
					{"Access to all methods", true, false},
					{"Unlimited assumptions per project", true, false},
					{"Generate executive reports", false, false},
					{"Invite team mates", false, false},
					{"Create multiple teams", false, false},
				},
				CTA: "Join beta list",
			},
			{
				CustomerType: "Startups",
				TagColor:     "cyan",
				Price:        "$99/mo",
				Highlights: []string{
					"Unlimited number of tests",
					"Unlimited number of assumptions",
				},
				Features: []Feature{
					{"Analytics dashboard", true, false},
					{"Access to all methods", true, false},
					{"Invite team mates", true, false},
					{"Generate executive reports", true, false},
					{"Limited to 1 project", false, false},
					{"Limited to 1 team", false, false},
				},
				CTA: "Join beta list",
			},
			{
				CustomerType: "Incubators",
				TagColor:     "yellow",
				Price:        "Custom",
				Highlights: []string{
					"Unlimited number of tests",
					"Unlimited number of assumptions",
				},
				Features: []Feature{
					{"Analytics dashboard", true, false},
					{"Access to all methods", true, false},
					{"Multiple projects", true, false},
					{"Multiple teams", true, false},
					{"Automated executive reports", false, true},
					{"Growth board", false, true},
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
