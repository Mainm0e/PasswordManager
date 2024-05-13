package data

import (
	"os/exec"
	"testing"
)

func TestCreateDataBase(t *testing.T) {
	err := CreateDataBase()
	if err != nil {
		t.Error("Error creating database: ", err)
	}
}

func TestRegisterAccount(t *testing.T) {
	err := RegisterAccount("exampleUsername10", "examplePassword10")
	if err != nil {
		t.Error(err)
	}

	err = RegisterAccount("exampleUsername11", "examplePassword11")
	if err != nil {
		t.Error(err)
	}

	err = RegisterAccount("exampleUsername12", "examplePassword13")
	if err != nil {
		t.Error(err)
	}

	// Test for duplicate username
	err = RegisterAccount("exampleUsername10", "examplePassword10")
	if err == nil {
		t.Error("Expected error for duplicate username")
	}
}

func TestLogin(t *testing.T) {
	// Test successful logins
	_, err := Login("exampleUsername10", "examplePassword10")
	if err != nil {
		t.Error(err)
	}

	_, err = Login("exampleUsername11", "examplePassword11")
	if err != nil {
		t.Error(err)
	}

	_, err = Login("exampleUsername12", "examplePassword13")
	if err != nil {
		t.Error(err)
	}

	// Test for non-existing username
	_, err = Login("exampleUsername13", "examplePassword13")
	if err == nil {
		t.Error("Expected error for non-existing username")
	}

	// Test for incorrect password
	_, err = Login("exampleUsername10", "examplePassword11")
	if err == nil {
		t.Error("Expected error for incorrect password")
	}
}

func TestDatabaseExistence(t *testing.T) {
	if IsDatabaseExit() {
		t.Log("Database exists")
		cmd := exec.Command("rm", "./accounts.db")
		err := cmd.Run()
		if err != nil {
			t.Error("Error removing database: ", err)
		}
	} else {
		t.Error("Database does not exist")
	}
}
