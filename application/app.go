package application

import (
	"database/sql"
	"fmt"
	"passwordmanager/PassGenerator"

	_ "github.com/mattn/go-sqlite3"
)

func App() {
	// Generate a password of length 12
	password := PassGenerator.GeneratePassword(20)
	fmt.Println("Generated Password:", password)
	createDatabase()
}

func createDatabase() {
	// Open SQLite database file. If it doesn't exist, it will be created.
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS passwords (
                        id INTEGER PRIMARY KEY,
                        password TEXT)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Insert some sample data
	password := "securepassword123"
	_, err = db.Exec("INSERT INTO passwords(password) VALUES (?)", password)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}

	fmt.Println("Data inserted successfully.")
}
