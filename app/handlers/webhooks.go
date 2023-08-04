package handlers

import (
	// "os"

	"github.com/batudal/derisk/app/config"
	// "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
)

func GithubWebhook(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jsonMap := make(map[string](interface{}))
		err := c.BodyParser(&jsonMap)
		if err != nil {
			return err
		}
		println(jsonMap)
		for key, value := range jsonMap {
			println("Key:", key, "Value:", value)
		}

		// bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
		// if err != nil {
		//     panic(err)
		// }
		return c.SendString("OK")
	}
}
