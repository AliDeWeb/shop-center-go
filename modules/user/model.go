package auth

import "github.com/kamva/mgm/v3"

type MUser struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"username" validate:"required"`
	Email            string `bson:"email" validate:"required,email" unique:"true"`
	Password         string `bson:"password" validate:"required,min=8"`
}
