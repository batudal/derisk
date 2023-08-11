package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WaitListSignup struct {
	Email        string             `bson:"email" validate:"required,email"`
	CustomerType string             `bson:"customer_type"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

type Feedback struct {
	Email        string             `bson:"email" validate:"required,email"`
	CustomerType string             `bson:"customer_type"`
	Message      string             `bson:"message" validate:"required,min=10,max=500"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

type EmailRequest struct {
	Customer  Customer
	EmailType string
}

type Customer struct {
	Email        string
	CustomerType string
}
