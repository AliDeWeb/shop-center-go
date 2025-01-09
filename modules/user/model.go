package auth

import (
	"time"
)

type MUser struct {
	Name      string    `bson:"username" validate:"required"`
	Email     string    `bson:"email" validate:"required,email" unique:"true"`
	Password  string    `bson:"password" validate:"required,min=8"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
