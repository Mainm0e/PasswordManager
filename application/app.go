package application

import (
	"fmt"
	"passwordmanager/PassGenerator"

	_ "github.com/mattn/go-sqlite3"
)

func App() {
	// Generate a password of length 12
	password := PassGenerator.GeneratePassword(20)
	fmt.Println("Generated Password:", password)
}
