package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	start := time.Now()
	password := "secret"
	hash, _ := HashPassword(password)
	fmt.Println("hash cost in ", time.Since(start))
	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)
	start = time.Now()
	match := CheckPasswordHash(password, hash)
	fmt.Println("check cost in ", time.Since(start))
	fmt.Println("Match:   ", match)
	
}
