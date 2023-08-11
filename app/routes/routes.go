package routes

import (
	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/handlers"
)

func Listen(cfg *config.Config) {
	cfg.App.Get("/", handlers.HandleIndexPage(cfg))
	cfg.App.Get("/about", handlers.HandleAboutPage(cfg))
	cfg.App.Get("/blog", handlers.HandleBlogPage(cfg))
	cfg.App.Get("/pricing", handlers.HandlePricingPage(cfg))
	cfg.App.Get("/contact", handlers.HandleContactPage(cfg))
	cfg.App.Get("/modals/wait-list", handlers.HandleWaitList(cfg))
	cfg.App.Get("/modals/feedback", handlers.HandleFeedback(cfg))
	cfg.App.Get("/modals/mobile-menu", handlers.HandleMobileMenu(cfg))
	cfg.App.Post("/join/wait-list", handlers.JoinWaitList(cfg))
	cfg.App.Post("/feedback", handlers.SendFeedback(cfg))
	cfg.App.Post("/webhooks/github", handlers.GithubWebhook(cfg))
	cfg.App.Static("/public", "./public")
	cfg.App.Listen(":80")
}
