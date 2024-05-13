package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDataBase(dataBasePath string) error {
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
        account_count INTEGER,
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

	// SQL statement to create the applicationdata table
	sqlStmtApplicationData := `
    CREATE TABLE IF NOT EXISTS applicationdata (
        id INTEGER PRIMARY KEY,
        user_id INTEGER,
        application_id INTEGER,
        username TEXT,
        password TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(application_id) REFERENCES applications(id)
    );
    `
	// Execute the SQL statement to create the applicationdata table
	_, err = db.Exec(sqlStmtApplicationData)
	if err != nil {
		return fmt.Errorf("error creating applicationdata table: %w", err)
	}
	fmt.Println("Created applicationdata table")

	return nil
}

func RegisterAccount(dataBasePath, username, password string) error {
	// Hash the password
	password = hashing(password)

	// Check if the database exists
	if !IsDatabaseExit(dataBasePath) {
		return fmt.Errorf("database does not exist")
	}

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// Check if the username already exists
	rows, err := db.Query("SELECT id FROM users WHERE username = ?", username)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		return fmt.Errorf("username already exists")
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating over query results: %w", err)
	}

	// SQL statement to insert data into the user table
	sqlStmt := `
    INSERT INTO users (username, password) VALUES (?, ?)
    `
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt, username, password)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}

func Login(dataBasePath, username, password string) (string, error) {
	// Hash the provided password
	hashedPassword := hashing(password)

	// Check if the database exists
	if !IsDatabaseExit(dataBasePath) {
		return "", fmt.Errorf("database does not exist")
	}

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		return "", fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to query the database
	rows, err := db.Query("SELECT id, username, password FROM users WHERE username = ?", username)
	if err != nil {
		return "", fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	// Check if any rows were returned
	if !rows.Next() {
		return "", fmt.Errorf("invalid username or password")
	}

	// If rows were returned, retrieve the hashed password from the database
	var id string
	var storedUsername string
	var storedPassword string
	if err := rows.Scan(&id, &storedUsername, &storedPassword); err != nil {
		return "", fmt.Errorf("error scanning row: %w", err)
	}

	// Compare the hashed password from the database with the hashed provided password
	if storedPassword != hashedPassword {
		return "", fmt.Errorf("invalid username or password")
	}

	// If the passwords match, the user is authenticated
	// You can return additional information about the user if needed
	return id, nil
}

func AddApplication(dataBasePath, userId, name, url string) error {
	// Check if the database exists
	if !IsDatabaseExit(dataBasePath) {
		return fmt.Errorf("database does not exist")
	}

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// Check if the user with the provided userId exists
	var userCount int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userId).Scan(&userCount)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	if userCount == 0 {
		return fmt.Errorf("user with the provided ID does not exist")
	}

	// Check if an application with the same name already exists for the user
	var appCount int
	err = db.QueryRow("SELECT COUNT(*) FROM applications WHERE user_id = ? AND name = ?", userId, name).Scan(&appCount)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	if appCount > 0 {
		return fmt.Errorf("an application with the same name already exists for the user")
	}

	// SQL statement to insert data into the applications table
	sqlStmt := `
    INSERT INTO applications (user_id, account_count, name, url) VALUES (?, ?, ?, ?)
    `
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt, userId, 0, name, url)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}

func AddApplicationData(dataBasePath, userId, applicationId, username, password string) error {
	// Check if the database exists
	if !IsDatabaseExit(dataBasePath) {
		return fmt.Errorf("database does not exist")
	}

	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// Check if the user with the provided userId exists
	var userCount int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userId).Scan(&userCount)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	if userCount == 0 {
		return fmt.Errorf("user with the provided ID does not exist")
	}

	// Check if the application with the provided applicationId exists
	var appCount int
	err = db.QueryRow("SELECT COUNT(*) FROM applications WHERE id = ?", applicationId).Scan(&appCount)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	if appCount == 0 {
		return fmt.Errorf("application with the provided ID does not exist")
	}

	// SQL statement to insert data into the applicationdata table
	sqlStmt := `
	INSERT INTO applicationdata (user_id, application_id, username, password) VALUES (?, ?, ?, ?)
	`
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt, userId, applicationId, username, password)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	// Update the account_count for the application
	_, err = db.Exec("UPDATE applications SET account_count = account_count + 1 WHERE id = ?", applicationId)
	if err != nil {
		return fmt.Errorf("error updating account count: %w", err)
	}

	return nil

}

func IsDatabaseExit(dataBasePath string) bool {
	// Check if the database file exists
	if _, err := os.Stat(dataBasePath); err == nil {
		// Database file exists
		return true
	}

	return false
}
