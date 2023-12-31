package main

import (
	"flag"
	"os"
	"time"

	"context"

	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/middleware"
	"github.com/batudal/derisk/app/routes"
	"github.com/batudal/derisk/app/schema"
	"github.com/batudal/derisk/app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/resendlabs/resend-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := setup()
	cfg.App.Use(logger.New())
	cfg.App.Use(recover.New())
	cfg.App.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
	}))
	cfg.App.Use(cors.New(cors.Config{
		AllowHeaders:  "HX-Request, HX-Trigger, HX-Trigger-Name, HX-Target, HX-Prompt",
		ExposeHeaders: "HX-Push, HX-Redirect, HX-Location, HX-Refresh, HX-Trigger, HX-Trigger-After-Swap, HX-Trigger-After-Settle",
	}))
	cfg.App.Use(helmet.New(helmet.Config{
		XSSProtection:             "1; mode=block",
		CrossOriginEmbedderPolicy: "unsafe-none",
		CrossOriginResourcePolicy: "cross-origin",
		CrossOriginOpenerPolicy:   "unsafe-none",
	}))
	cfg.App.Use(middleware.AddCacheHeaders(cfg))
	go services.ListenEmailRequests(&cfg)
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
		App:           app,
		Mc:            mongodb_client,
		Rs:            resend_client,
		EmailRequests: make(chan schema.EmailRequest),
		LastModified:  time.Now().Format(time.RFC1123),
	}
	return cfg
}
