package handlers

import (
	"os"

	"github.com/batudal/derisk/app/config"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
)

func GithubWebhook(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		println("Headers:", headers)
		println("Signature:", headers["X-Hub-Signature-256"])
		jsonMap := make(map[string](interface{}))
		err := c.BodyParser(&jsonMap)
		if err != nil {
			return err
		}
		pusher := jsonMap["pusher"].(map[string]interface{})
		name := pusher["name"].(string)
		bot, err := tgbot.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
		if err != nil {
			panic(err)
		}
		msg := tgbot.NewMessageToChannel(os.Getenv("TELEGRAM_CHANNEL"), name+" just pushed to "+jsonMap["repository"].(map[string]interface{})["name"].(string))
		msg.ParseMode = "markdown"
		bot.Send(msg)
		return c.SendString("OK")
	}
}
