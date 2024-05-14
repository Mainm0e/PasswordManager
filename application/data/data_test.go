package data

import (
	"os/exec"
	"testing"
)

var (
	// Test database path
	testDataBasePath = "./test.db"
)

func TestCreateDataBase(t *testing.T) {
	err := CreateDataBase(testDataBasePath)
	if err != nil {
		t.Error("Error creating database: ", err)
	}
}
func TestRegisterAccount(t *testing.T) {
	db, err := OpenDatabaseConnection(testDataBasePath)
	if err != nil {
		t.Error(err)
	}

	err = RegisterAccount(db, "exampleUsername10", "examplePassword10")
	if err != nil {
		t.Error(err)
	}

	err = RegisterAccount(db, "exampleUsername11", "examplePassword11")
	if err != nil {
		t.Error(err)
	}

	err = RegisterAccount(db, "exampleUsername12", "examplePassword13")
	if err != nil {
		t.Error(err)
	}

	// Test for duplicate username
	err = RegisterAccount(db, "exampleUsername10", "examplePassword10")
	if err == nil {
		t.Error("Expected error for duplicate username")
	}
}

func TestLogin(t *testing.T) {
	db, err := OpenDatabaseConnection(testDataBasePath)
	if err != nil {
		t.Error(err)
	}
	// Test successful logins
	_, err = Login(db, "exampleUsername10", "examplePassword10")
	if err != nil {
		t.Error(err)
	}

	_, err = Login(db, "exampleUsername11", "examplePassword11")
	if err != nil {
		t.Error(err)
	}

	_, err = Login(db, "exampleUsername12", "examplePassword13")
	if err != nil {
		t.Error(err)
	}

	// Test for non-existing username
	_, err = Login(db, "exampleUsername13", "examplePassword13")
	if err == nil {
		t.Error("Expected error for non-existing username")
	}

	// Test for incorrect password
	_, err = Login(db, "exampleUsername10", "examplePassword11")
	if err == nil {
		t.Error("Expected error for incorrect password")
	}
}

func TestAddApplication(t *testing.T) {
	db, err := OpenDatabaseConnection(testDataBasePath)
	if err != nil {
		t.Error(err)
	}
	// Test successful application addition
	err = AddApplication(db, "1", "exampleApp1", "exampleURL1")
	if err != nil {
		t.Error(err)
	}

	err = AddApplication(db, "2", "exampleApp2", "exampleURL2")
	if err != nil {
		t.Error(err)
	}

	err = AddApplication(db, "2", "exampleApp3", "exampleURL3")
	if err != nil {
		t.Error(err)
	}

	// Test for non-existing username
	err = AddApplication(db, "5", "exampleApp4", "exampleURL4")
	if err == nil {
		t.Error("Expected error for non-existing username")
	}
}

func TestAddApplicationData(t *testing.T) {
	db, err := OpenDatabaseConnection(testDataBasePath)
	if err != nil {
		t.Error(err)
	}
	// Test successful application data addition
	err = AddApplicationData(db, "1", "1", "exampleUsername1", "examplePassword1")
	if err != nil {
		t.Error(err)
	}

	err = AddApplicationData(db, "1", "2", "exampleUsername2", "examplePassword2")
	if err != nil {
		t.Error(err)
	}

	err = AddApplicationData(db, "2", "3", "exampleUsername3", "examplePassword3")
	if err != nil {
		t.Error(err)
	}

	// test adding data for an existing application
	err = AddApplicationData(db, "1", "1", "exampleUsername4", "examplePassword4")
	if err != nil {
		t.Error("Expected error for existing application")
	}

	// Test for non-existing username
	err = AddApplicationData(db, "5", "4", "exampleUsername4", "examplePassword4")
	if err == nil {
		t.Error("Expected error for non-existing username")
	}
}

// test getting datas
func TestGetApplications(t *testing.T) {
	db, err := OpenDatabaseConnection(testDataBasePath)
	if err != nil {
		t.Error(err)
	}

	apps, err := GetApplications(db, "1")
	if err != nil {
		t.Error(err)
	}

	if len(apps) != 1 {
		t.Error("Expected 1 application, got ", len(apps))
	}

	apps, err = GetApplications(db, "2")
	if err != nil {
		t.Error(err)
	}

	if len(apps) != 2 {
		t.Error("Expected 2 applications, got ", len(apps))
	}

	apps, err = GetApplications(db, "3")
	if err != nil {
		t.Error(err)
	}

	if len(apps) != 0 {
		t.Error("Expected 0 applications, got ", len(apps))
	}
}

func TestGetApplicationData(t *testing.T) {
	db, err := OpenDatabaseConnection(testDataBasePath)
	if err != nil {
		t.Error(err)
	}

	data, err := GetApplicationData(db, "1", "2")
	if err != nil {
		t.Error(err)
	}

	if len(data) != 1 {
		t.Error("Expected 1 data, got ", len(data))
	}

	data, err = GetApplicationData(db, "2", "3")
	if err != nil {
		t.Error(err)
	}

	if len(data) != 1 {
		t.Error("Expected 1 data, got ", len(data))
	}

	data, err = GetApplicationData(db, "3", "4")
	if err == nil {
		t.Error("Expected error for non-existing application")
	}

	if len(data) != 0 {
		t.Error("Expected 0 data, got ", len(data))
	}
}

func TestDatabaseExistence(t *testing.T) {
	if IsDatabaseExit(testDataBasePath) {
		t.Log("Database exists")
		cmd := exec.Command("rm", testDataBasePath)
		err := cmd.Run()
		if err != nil {
			t.Error("Error removing database: ", err)
		}
	} else {
		t.Error("Database does not exist")
	}
}
