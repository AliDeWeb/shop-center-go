package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MUser struct {
	_id       primitive.ObjectID `bson:"-"`
	Name      string             `bson:"username" validate:"required"`
	Email     string             `bson:"email" validate:"required,email" unique:"true"`
	Password  string             `bson:"password" validate:"required,min=8"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Token     string             `bson:"token"`
}
