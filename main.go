package main

import (
	"fmt"
	"log"
	"passwordmanager/application"
	"passwordmanager/application/data"
)

func main() {
	db, err := data.OpenDatabaseConnection("passwordmanager.db")
	if err != nil {
		log.Fatal(err)
	}

	// Asking Usernames and Passwords
	fmt.Println("Enter your username:")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Enter your password:")
	var password string
	fmt.Scanln(&password)

	user_id, err := data.Login(db, username, password)

	if err != nil {
		fmt.Println("Login failed:", err)
		fmt.Println("Would you like to register a new account? (y/n)")
		var choice string
		fmt.Scanln(&choice)
		if choice == "y" {
			fmt.Println("Enter your new username:")
			var new_username string
			fmt.Scanln(&new_username)

			fmt.Println("Enter your new password:")
			var new_password string
			fmt.Scanln(&new_password)

			err := data.RegisterAccount(db, new_username, new_password)
			if err != nil {
				fmt.Println("Account registration failed:", err)
				return
			}
			fmt.Println("Account registered successfully!")
			user_id, err = data.Login(db, new_username, new_password)
			if err != nil {
				fmt.Println("Login failed:", err)
				return
			}
		} else {
			fmt.Println("Goodbye!")
			return
		}
	}
	fmt.Println("Welcome", username, "!")

	fmt.Println("Do you want to add a new Password or read your Passwords? (add/read)")
	var choice string
	fmt.Scanln(&choice)

	if choice == "add" {
		application.AddNewPassword(db, user_id)
	}
	if choice == "read" {
		application.ReadPasswords(db, user_id)
	}
	db.Close()
}
