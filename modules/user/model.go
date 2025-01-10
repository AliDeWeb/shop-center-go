package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"username" binding:"required,min=2,max=50"`
	Email     string             `bson:"email" binding:"required,email" unique:"true"`
	Password  string             `bson:"password" binding:"required,min=8,max=100"`
	Token     string             `bson:"token"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
