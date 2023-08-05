package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"os"

	"github.com/batudal/derisk/app/config"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
)

func GithubWebhook(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		hash := hmac.New(sha256.New, c.Body())
		secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
		hash.Write([]byte(secret))
		if !hmac.Equal(hash.Sum(nil), []byte(headers["X-Hub-Signature-256"])) {
			return c.SendStatus(403)
		}
		// temp
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
