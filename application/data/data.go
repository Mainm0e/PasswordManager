package data

import (
	"database/sql"
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

func createDataBase() {
	// Check if the database file exists
	if _, err := os.Stat(dataBasePath); err == nil {
		// Database file exists, no need to create new database or tables
		return
	} else if !os.IsNotExist(err) {
		// Error occurred while checking for existence, log and exit
		log.Fatal("Error checking database file existence:", err)
	}
	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", dataBasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to create the passwords table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY,
		username TEXT,
		password TEXT
	);
	`
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	sqlStmtApplications := `
	CREATE TABLE IF NOT EXISTS applications (
		id INTEGER PRIMARY KEY,
		user_id INTEGER,
		name TEXT,
		FOREIGN KEY(user_id) REFERENCES user(id)
	);
	`
	// Execute the SQL statement to create the applications table
	_, err = db.Exec(sqlStmtApplications)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmtApplications)
	}

	// SQL statement to create the applications table
	sqlStmtData := `
	CREATE TABLE IF NOT EXISTS data (
		id INTEGER PRIMARY KEY,
		user_id INTEGER,
		application_id INTEGER,
		username TEXT,
		password TEXT,
		order INTEGER,
		FOREIGN KEY(user_id) REFERENCES user(id)
		FOREIGN KEY(application_id) REFERENCES applications(id)
	);
	`
	// Execute the SQL statement to create the applications table
	_, err = db.Exec(sqlStmtData)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmtData)
	}
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
	INSERT INTO user (username, password) VALUES (?, ?)
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
	rows, err := db.Query("SELECT id, username, password FROM user")
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
