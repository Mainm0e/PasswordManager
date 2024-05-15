package application

import (
	"database/sql"
	"fmt"
	"passwordmanager/PassGenerator"
	"passwordmanager/application/data"
	"strconv"
)

// Todo: Encrypt new Passwords before saving them to the database
// Key: 16 bytes long and Key is generated based on the user's hashed password
func AddNewPassword(db *sql.DB, user_id string) {
	var applicationName, applicationURL, username string
	var passwordLength int
	var applicationID int
	for {
		choice := getInput("Do you want to Add new application (y/n)")
		if choice == "y" {
			// Prompt the user for application details
			applicationName = getInput("Enter the application name:")
			applicationURL = getInput("Enter the application URL:")
			// Save the application
			app_id, err := data.AddApplication(db, user_id, applicationName, applicationURL)
			if err != nil {
				fmt.Println("Error adding application:", err)
				return
			}
			applicationID = app_id
			fmt.Println("Application saved successfully!")
			break
		} else {
			fmt.Println("Here are the applications you have:")
			apps, err := data.GetApplications(db, user_id)
			if err != nil {
				fmt.Println("Error getting applications:", err)
				return
			}
			if len(apps) == 0 {
				fmt.Println("No applications found.")
				continue
			}
			for _, app := range apps {
				fmt.Println("Name:", app.Name, "URL", app.URL, "Application ID", app.ID)
			}
			for {
				choice := getIntInput("Enter the application ID you want to add password to:")
				//check if the choice is equal to the application ID
				for _, app := range apps {
					if choice == app.ID {
						fmt.Println("You have selected", app.Name, "Application")
						applicationID = choice
						break
					}
				}
				if applicationID != 0 {
					break
				} else {
					fmt.Println("Invalid application ID")
					choice2 := getInput("Do you want to try again? (y/n)")
					if choice2 != "y" {
						continue
					}
				}
			}
			if applicationID != 0 {
				break
			}
		}
	}
	// Prompt the user for username
	username = getInput("Enter the username:")
	// Prompt the user for password length
	passwordLength = getIntInput("Enter the password length:")

	for {
		// Generate a password
		password := PassGenerator.GeneratePassword(passwordLength)
		fmt.Println("Generated password:", password)

		// Ask if the user wants to save the password
		choice := getInput("Do you want to save this password? (y/n)")

		if choice == "y" {
			// Save the password
			err := data.AddApplicationData(db, user_id, strconv.Itoa(applicationID), username, password)
			if err != nil {
				fmt.Println("Error saving password:", err)
				continue
			}
			fmt.Println("Password saved successfully!")
			break
		} else {
			fmt.Println("Password not saved.")
		}

		// Ask if the user wants to generate another password
		choice = getInput("Do you want to generate another password? (y/n)")
		if choice != "y" {
			break // Exit the loop if the user doesn't want to generate another password
		}
	}

	fmt.Println("Goodbye!")
}

func getInput(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func getIntInput(prompt string) int {
	input := getInput(prompt)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return getIntInput(prompt)
	}
	return num
}

// GetApplicationData is function that user can read application account data from database
func ReadPasswords(db *sql.DB, user_id string) {

	// Get the applications
	apps, err := data.GetApplications(db, user_id)
	if err != nil {
		fmt.Println("Error getting applications:", err)
		return
	}
	if len(apps) == 0 {
		fmt.Println("No applications found.")
		return
	}
	for _, app := range apps {
		fmt.Println("\nName:", app.Name, "\nURL:", app.URL, "\nApplication ID:", app.ID)
	}

	// Prompt the user for the application ID
	applicationID := getIntInput("\nEnter the application ID you want to get data for:")

	// Get the application data
	appData, err := data.GetApplicationData(db, user_id, strconv.Itoa(applicationID))
	if err != nil {
		fmt.Println("Error getting application data:", err)
		return
	}
	if len(appData) == 0 {
		fmt.Println("No data found for the application.")
		return
	}
	for index, data := range appData {
		fmt.Println("\nNo.", index, "\nUsername:", data.Username, "\nPassword:", data.Password)
	}
	fmt.Println("")

}
