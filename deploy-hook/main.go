package main

import (
	"crypto/tls"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//temp
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		hook, _ := github.New(github.Options.Secret(os.Getenv("GITHUB_WEBHOOK_SECRET")))
		httpRequest := new(http.Request)
		err := fasthttpadaptor.ConvertRequest(c.Context(), httpRequest, true)
		if err != nil {
			println(err.Error())
			c.SendStatus(400)
		}
		payload, err := hook.Parse(httpRequest, github.PushEvent)
		if err != nil {
			println(err.Error())
			c.SendStatus(400)
		}
		var out []byte
		switch payload.(type) {
		case github.PushPayload:
			cmd := exec.Command("./update_services.sh -y")
			out, err = cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
		}
		return c.SendString(string(out))
	})
	cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":37373", config)
	if err != nil {
		panic(err)
	}
	log.Fatal(app.Listener(ln))
}
