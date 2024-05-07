package data

import (
	"os/exec"
	"testing"
)

func TestData(t *testing.T) {
	// Create a database
	//createDataBase()
	// check if the database is created

	// try to insert data into the database
	err := insertData()
	if err != nil {
		t.Error("Error inserting data")
	}
	// delete password.db

	cmd := exec.Command("rm", "passwords.db")
	err = cmd.Run()

	if err != nil {
		t.Error("Error inserting data")
	}
}
