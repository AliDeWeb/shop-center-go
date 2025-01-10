package auth

import (
	"sync"
	"time"

	"github.com/alideweb/shop-center-go/db"
	"github.com/alideweb/shop-center-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type tokens struct {
	AccessToken  string
	RefreshToken string
}

func SRegisterUser(user *MUser) (*mongo.InsertOneResult, *MUser, *tokens, error) {
	accessToken := utils.GenerateJWT(user.Email, 15)
	refreshToken := utils.GenerateJWT(user.Email, 60*24*7)

	var wg sync.WaitGroup
	wg.Add(2)
	passwordChan := make(chan string)
	refreshTokenChan := make(chan string)
	defer close(passwordChan)
	defer close(refreshTokenChan)

	go func() {
		defer wg.Done()
		utils.Hash(user.Password, passwordChan)
	}()
	go func() {
		defer wg.Done()
		utils.Hash(refreshToken, refreshTokenChan)
	}()
	wg.Wait()

	user.Password = <-passwordChan
	user.Token = <-refreshTokenChan
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, data, err := db.InsertOne("user", user)

	tokens := &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return result, data, tokens, err
}
