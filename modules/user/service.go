package auth

import (
	"github.com/alideweb/shop-center-go/db"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SRegisterUser(user *MUser) (*mongo.InsertOneResult, *MUser, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	user.Password = string(hashedPassword)

	result, data, err := db.InsertOne("user", user)

	return result, data, err
}
