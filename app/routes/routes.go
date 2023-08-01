package routes

import (
	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/handlers"
)

func Listen(cfg *config.Config) {
	cfg.App.Get("/", handlers.Index(cfg))
	cfg.App.Static("/public", "./public")
	cfg.App.Listen(":80")
}
