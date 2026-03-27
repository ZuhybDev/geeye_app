package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(pass []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	if err != nil {
		// NEVER use log.Fatal in a handler! It kills the server process.
		fmt.Printf("Hashing password Error: %v", err)
	}
	return string(hash), nil
}

func VerifyPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
