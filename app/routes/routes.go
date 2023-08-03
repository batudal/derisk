package routes

import (
	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/handlers"
)

func Listen(cfg *config.Config) {
	cfg.App.Get("/", handlers.Index(cfg))
	cfg.App.Get("/modals/beta-list", handlers.HandleBetaList(cfg))
	cfg.App.Get("/modals/feedback", handlers.HandleFeedback(cfg))
	cfg.App.Post("/join/beta-list", handlers.JoinBetaList(cfg))
	cfg.App.Post("/feedback", handlers.SendFeedback(cfg))
	cfg.App.Static("/public", "./public")
	cfg.App.Listen(":80")
}
