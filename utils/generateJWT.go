package utils

import (
	"fmt"
	"time"

	"github.com/alideweb/shop-center-go/config"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(data string, expiresInHour uint) string {
	claims := jwt.MapClaims{
		data:  data,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(expiresInHour) * time.Hour).Unix(),
	}

	secretKey := []byte(config.ServerEnvsConfig.JwtSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing the token:", err)
		return ""
	}

	return signedToken
}
