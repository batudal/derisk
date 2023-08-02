package handlers

import (
	"time"

	"github.com/batudal/derisk/app/config"
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
	Email     string             `bson:"email"`
	Message   string             `bson:"message"`
	CreatedAt primitive.DateTime `bson:"created_at"`
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
				return err
			}
			return c.Render("components/modals/beta-list-success", fiber.Map{})
		}
		return err
	}
}
