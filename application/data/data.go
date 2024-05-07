package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Data() {
	// Create a database
	createDataBase()
	// Insert data into the database
	insertData()
	// Retrieve data from the database
	retrieveData()
}

func createDataBase() {
	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", "./passwords.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to create the passwords table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS passwords (
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
}

func insertData() error {
	// Open a SQLite database connection
	db, err := sql.Open("sqlite3", "./passwords.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close() // Make sure to close the database connection when the function ends

	// SQL statement to insert data into the passwords table
	sqlStmt := `
	INSERT INTO passwords (username, password) VALUES ('user1', 'password1');
	`
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil

}

func retrieveData() {

}
