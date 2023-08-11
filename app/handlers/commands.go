package handlers

import (
	"time"

	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/schema"
	"github.com/batudal/derisk/app/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func JoinBetaList(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		validate := validator.New()
		err := validate.Struct(schema.BetaListSignup{
			Email: c.FormValue("email"),
		})
		if err != nil {
			return c.Render("components/modals/beta-list", fiber.Map{
				"Error": "Please enter a valid email address.",
			})
		}
		var signup schema.BetaListSignup
		coll := cfg.Mc.Database("mvp").Collection("beta-list-signups")
		err = coll.FindOne(c.Context(), bson.M{"email": c.FormValue("email")}).Decode(&signup)
		if err != nil && err == mongo.ErrNoDocuments {
			_, err := coll.InsertOne(c.Context(), schema.BetaListSignup{
				Email:        c.FormValue("email"),
				CustomerType: c.FormValue("customer_type"),
				CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
			})
			if err != nil {
				return c.Render("components/modals/beta-list", fiber.Map{
					"Error": "Something went wrong. Please try again later.",
				})
			}
			_ = services.RequestEmail(cfg, &schema.Customer{
				Email:        c.FormValue("email"),
				CustomerType: c.FormValue("customer_type"),
			}, "join-beta-list")
			return c.Render("components/modals/beta-list-success", fiber.Map{
				"Email":        c.FormValue("email"),
				"CustomerType": c.FormValue("customer_type"),
				"LeftFeedback": c.Query("left_feedback") == "true",
			})
		}
		return c.Render("components/modals/beta-list", fiber.Map{
			"Error": c.FormValue("email") + " is already registered. Please try again with a different email.",
		})
	}
}

func SendFeedback(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		validate := validator.New()
		err := validate.Struct(schema.Feedback{
			Email:   c.FormValue("email"),
			Message: c.FormValue("message"),
		})
		var email_error string
		var message_error string
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				if err.Field() == "Email" {
					email_error = "Please enter a valid email address."
				}
				if err.Field() == "Message" {
					message_error = "Please enter a message between 10-500 characters."
				}
			}
			return c.Render("components/modals/feedback", fiber.Map{
				"Email":        c.FormValue("email"),
				"CustomerType": c.FormValue("customer_type"),
				"EmailError":   email_error,
				"MessageError": message_error,
			})
		}
		coll := cfg.Mc.Database("mvp").Collection("feedback")
		_, err = coll.InsertOne(c.Context(), schema.Feedback{
			Email:        c.FormValue("email"),
			CustomerType: c.FormValue("customer_type"),
			Message:      c.FormValue("message"),
			CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
		})
		if err != nil {
			return c.Render("components/modals/feedback", fiber.Map{
				"Error": "Something went wrong. Please try again later.",
			})
		}
		_ = services.RequestEmail(cfg, &schema.Customer{
			Email:        c.FormValue("email"),
			CustomerType: c.FormValue("customer_type"),
		}, "feedback")
		subscribed := false
		signup_coll := cfg.Mc.Database("mvp").Collection("beta-list-signups")
		err = signup_coll.FindOne(c.Context(), bson.M{"email": c.FormValue("email")}).Decode(&schema.BetaListSignup{})
		if err == nil {
			subscribed = true
		}
		return c.Render("components/modals/feedback-success", fiber.Map{
			"Email":        c.FormValue("email"),
			"CustomerType": c.FormValue("customer_type"),
			"Subscribed":   subscribed,
		})
	}
}
