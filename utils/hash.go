package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	log.Println("Inside HashPassword")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Password hashing failed : %v", err)
		return "", err
	}
	log.Printf("hashed password : %v", hash)
	return string(hash), nil
}
