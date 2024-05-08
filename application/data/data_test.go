package data

import (
	"testing"
)

func TestData(t *testing.T) {
	// Create a database
	createDataBase()

	// check if the database is created

	// try to insert data into the database
	err := registerAccount("exampleUsername10", "examplePassword10")
	if err != nil {
		t.Error("Error inserting data")
	}
	err = registerAccount("exampleUsername11", "examplePassword11")
	if err != nil {
		t.Error("Error inserting data")
	}
	err = registerAccount("exampleUsername12", "examplePassword13")
	if err != nil {
		t.Error("Error inserting data")
	}
}
