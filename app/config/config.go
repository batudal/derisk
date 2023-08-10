package config

import (
	"github.com/batudal/derisk/app/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/resendlabs/resend-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	App           *fiber.App
	Mc            *mongo.Client
	Rs            *resend.Client
	EmailRequests chan schema.EmailRequest
	LastModified  string
}
