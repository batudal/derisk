package main

import (
	"flag"
	"os"

	"context"

	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/resendlabs/resend-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// todo: add limiter for spam
	cfg := setup()
	cfg.App.Use(logger.New())
	cfg.App.Use(cors.New(cors.Config{
		AllowHeaders:  "HX-Request, HX-Trigger, HX-Trigger-Name, HX-Target, HX-Prompt",
		ExposeHeaders: "HX-Push, HX-Redirect, HX-Location, HX-Refresh, HX-Trigger, HX-Trigger-After-Swap, HX-Trigger-After-Settle",
	}))
	routes.Listen(&cfg)
}

func setup() config.Config {
	dev := flag.Bool("dev", false, "development mode")
	flag.Parse()
	if *dev {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}
	uri := os.Getenv("MONGODB_URI")
	mongodb_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	engine := html.New("./views", ".html")
	if *dev {
		engine.Reload(true)
	}
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	resend_client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	cfg := config.Config{
		App: app,
		Mc:  mongodb_client,
		Rs:  resend_client,
	}
	return cfg
}
