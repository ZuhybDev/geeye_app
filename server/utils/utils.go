package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// generate
func GenerateToken(id, name, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
	})

	tkn, err := token.SignedString(os.Getenv("WT_SECRET"))

	if err != nil {
		return "Error at signing jwt", err
	}

	return tkn, nil
}

// verify it

func VerifyToken(tknString string) (bool, error) {
	token, err := jwt.Parse(tknString, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("WT_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
