package handlers

import (
	"time"

	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BetaListSignup struct {
	Email        string             `bson:"email"`
	CustomerType string             `bson:"customer_type"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

type Feedback struct {
	Email        string             `bson:"email"`
	CustomerType string             `bson:"customer_type"`
	Message      string             `bson:"message"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

func JoinBetaList(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var signup BetaListSignup
		coll := cfg.Mc.Database("mvp").Collection("beta-list-signups")
		err := coll.FindOne(c.Context(), bson.M{"email": c.FormValue("email")}).Decode(&signup)
		if err != nil && err == mongo.ErrNoDocuments {
			_, err := coll.InsertOne(c.Context(), BetaListSignup{
				Email:        c.FormValue("email"),
				CustomerType: c.FormValue("customer_type"),
				CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
			})
			if err != nil {
				return c.Render("components/modals/beta-list", fiber.Map{
					"Error": "Something went wrong. Please try again later.",
				})
			}
			err = services.JoinBetaListEmail(cfg, c.FormValue("email"), c.FormValue("customer_type"))
			if err != nil {
				return c.Render("components/modals/beta-list", fiber.Map{
					"Error": err.Error(),
				})
			}
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
		coll := cfg.Mc.Database("mvp").Collection("feedback")
		_, err := coll.InsertOne(c.Context(), Feedback{
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
		subscribed := false
		signup_coll := cfg.Mc.Database("mvp").Collection("beta-list-signups")
		err = signup_coll.FindOne(c.Context(), bson.M{"email": c.FormValue("email")}).Decode(&BetaListSignup{})
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
