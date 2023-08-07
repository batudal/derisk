package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/resendlabs/resend-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	App          *fiber.App
	Mc           *mongo.Client
	Rs           *resend.Client
	LastModified string
}
