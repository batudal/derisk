package handlers

import (
	"net/http"
	"os"

	"github.com/batudal/derisk/app/config"
	"github.com/go-playground/webhooks/v6/github"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func GithubWebhook(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hook, _ := github.New(github.Options.Secret(os.Getenv("GITHUB_WEBHOOK_SECRET")))
		httpRequest := new(http.Request)
		err := fasthttpadaptor.ConvertRequest(c.Context(), httpRequest, true)
		if err != nil {
			println(err.Error())
			c.SendStatus(400)
		}
		payload, err := hook.Parse(httpRequest, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			println(err.Error())
			c.SendStatus(400)
		}
		switch payload.(type) {
		case github.PushPayload:
			pusher_name := payload.(github.PushPayload).Pusher.Name
			SendMessage("New push from " + pusher_name)
		}
		return c.SendString("OK")
	}
}

func SendMessage(message string) {
	bot, err := tgbot.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		panic(err)
	}
	msg := tgbot.NewMessageToChannel(os.Getenv("TELEGRAM_CHANNEL"), message)
	msg.ParseMode = "markdown"
	bot.Send(msg)
}
