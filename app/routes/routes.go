package routes

import (
	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/handlers"
)

func Listen(cfg *config.Config) {
	cfg.App.Get("/", handlers.Index(cfg))
	cfg.App.Get("/modals/beta-list", handlers.BetaList(cfg))
	cfg.App.Post("/join/beta-list", handlers.JoinBetaList(cfg))
	cfg.App.Static("/public", "./public")
	cfg.App.Listen("0.0.0.0:80")
}
