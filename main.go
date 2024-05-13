package main

import (
	"fmt"
	"passwordmanager/PassGenerator"
	"passwordmanager/application/data"
)

func main() {
	data.Data()
	// Asking Usernames and Passwords
	fmt.Println("Enter your username:")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Enter your password:")
	var password string
	fmt.Scanln(&password)

	// Now you can do something with the username and password entered by the user
	// For now, let's just print them
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

	fmt.Println("Do you want to generate a password? (y/n)")
	var choice string
	fmt.Scanln(&choice)

	if choice == "y" {
		// Generate a password of length 12
		password := PassGenerator.GeneratePassword(12)
		fmt.Println("Generated Password:", password)
		fmt.Println("Press Ctrl+C to exit")
		for {
		}
	} else {
		fmt.Println("Are you looking for your password? (y/n)")
		var choice string
		fmt.Scanln(&choice)
		if choice == "y" {
			fmt.Println("Sorry, I can't help you with that.")
		} else {
			fmt.Println("Goodbye!")
		}
	}
}
