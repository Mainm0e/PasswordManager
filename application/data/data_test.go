package data

import (
	"testing"
)

func TestData(t *testing.T) {
	// Create a database
	createDataBase()
	// check if the database is created

	// try to insert data into the database
	err := registerAccount("exampleUsername", "examplePassword")
	if err != nil {
		t.Error("Error inserting data")
	}
}
