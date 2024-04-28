package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	symbols          = "!@#$&-_"
)

func generatePassword(length int) string {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Combine all possible characters
	allCharacters := lowercaseLetters + uppercaseLetters + digits + symbols

	// Initialize password slice
	password := make([]byte, length)

	// Fill password with random characters
	for i := 0; i < length; i++ {
		password[i] = allCharacters[rand.Intn(len(allCharacters))]
	}

	return string(password)
}

func main() {
	// Generate a password of length 12
	password := generatePassword(20)
	fmt.Println("Generated Password:", password)
}
