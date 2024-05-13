package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dataBasePath = "./accounts.db"
)

func Data() {
	// Create a database
	createDataBase()
	// Insert data into the database
	registerAccount("exampleUsername", "examplePassword")
	// Retrieve data from the database
	retrieveData()
}

func createDataBase() error {
	// Check if the database file exists
	if _, err := os.Stat(dataBasePath); err == nil {
		// Database file exists, no need to create new database or tables
		return nil
	} else if !os.IsNotExist(err) {
		// Error occurred while checking for existence, return the error
		return fmt.Errorf("error checking database file existence: %w", err)
	}

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to create the users table
	sqlStmtUsers := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT,
        password TEXT
    );
    `
	// Execute the SQL statement to create the users table
	_, err = db.Exec(sqlStmtUsers)
	if err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}
	fmt.Println("Created users table")

	// SQL statement to create the applications table
	sqlStmtApplications := `
    CREATE TABLE IF NOT EXISTS applications (
        id INTEGER PRIMARY KEY,
		user_id INTEGER,
		index INTEGER,
        name TEXT,
		url TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
    );
    `
	// Execute the SQL statement to create the applications table
	_, err = db.Exec(sqlStmtApplications)
	if err != nil {
		return fmt.Errorf("error creating applications table: %w", err)
	}
	fmt.Println("Created applications table")

	// SQL statement to create the accountdata table
	sqlStmtAccountdata := `
	CREATE TABLE IF NOT EXISTS accountdata (
		id INTEGER PRIMARY KEY,
		application_id INTEGER,
		user_id INTEGER,
		username TEXT,
		password TEXT,
		date_created TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(application_id) REFERENCES applications(id)
	);
	`
	// Execute the SQL statement to create the applicationdata table
	_, err = db.Exec(sqlStmtAccountdata)
	if err != nil {
		return fmt.Errorf("error creating applicationdata table: %w", err)
	}
	fmt.Println("Created applicationdata table")

	return nil
}

func registerAccount(username string, password string) error {

	password = hashing(password)

	if !isDatabaseExit() {
		createDataBase()
	}

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to insert data into the user table
	sqlStmt := `
	INSERT INTO users (username, password) VALUES (?, ?)
	`
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt, username, password)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}

func retrieveData() {

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to query the database
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop through the rows
	for rows.Next() {
		var id int
		var username string
		var password string
		err = rows.Scan(&id, &username, &password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, username, password)
	}

}

func isDatabaseExit() bool {
	// Check if the database file exists
	if _, err := os.Stat(dataBasePath); err == nil {
		// Database file exists
		return true
	}

	return false
}
