package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(s string, ch chan<- string) {
	go func() {
		sha256Hash := sha256.New()
		sha256Hash.Write([]byte(s))
		sha256HashedData := sha256Hash.Sum(nil)
		sha256HashedString := hex.EncodeToString(sha256HashedData)

		bcryptHashedData, err := bcrypt.GenerateFromPassword([]byte(sha256HashedString), 12)
		if err != nil {
			fmt.Println(err)
		}

		ch <- string(bcryptHashedData)
	}()
}

func CompareHash(hashed, plain string, ch chan<- bool) {
	go func() {
		sha256Hash := sha256.New()
		sha256Hash.Write([]byte(plain))
		sha256HashedData := sha256Hash.Sum(nil)
		sha256HashedString := hex.EncodeToString(sha256HashedData)

		err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(sha256HashedString))
		ch <- err == nil
	}()
}
