package application

import (
	"os/exec"
	"passwordmanager/application/data"
	"passwordmanager/application/tools"
	"strconv"
	"testing"
)

var (
	database = "./test.db"
	// Test User
	username = "testUser"
	password = "testPassword"

	// Test case for application
	applicationName      = "Test Application"
	applicationURL       = "http://test.com"
	application_username = "testUsername"
	application_password = "testPassword"
)

func TestSetupTestCase(t *testing.T) {
	t.Log("Test setup")

	// create a new database for testing
	err := data.CreateDataBase(database)
	if err != nil {
		t.Error("Error creating database: ", err)
	}

	// add a new user
	db, err := data.OpenDatabaseConnection(database)
	if err != nil {
		t.Error(err)
	}

	err = data.RegisterAccount(db, username, password)
	if err != nil {
		t.Error(err)
	}

	t.Log("login test user")
	user_id, err := data.Login(db, username, password)

	if err != nil {
		t.Error(err)
	}

	t.Log("Test setup complete")

	// Add a test application for testing and applicationdata
	application_id, err := data.AddApplication(db, user_id, applicationName, applicationURL)
	if err != nil {
		t.Error(err)
	}

	key, err := tools.GenerateKey(password, 32)
	if err != nil {
		t.Error(err)
	}

	// Encrypt the username and password
	encryptedUsername, err := tools.Encrypt(application_username, key)
	if err != nil {
		t.Error(err)
	}

	encryptedPassword, err := tools.Encrypt(application_password, key)
	if err != nil {
		t.Error(err)
	}

	// Add a test password for testing and applicationdata
	err = data.AddApplicationData(db, user_id, strconv.Itoa(application_id), encryptedUsername, encryptedPassword)
	if err != nil {
		t.Error(err)
	}

	//close the database connection
	db.Close()
}

func TestGetApplicationData(t *testing.T) {
	db, err := data.OpenDatabaseConnection(database)
	if err != nil {
		t.Error(err)
	}

	user_id, err := data.Login(db, username, password)
	if err != nil {
		t.Error(err)
	}

	key, err := tools.GenerateKey(password, 32)
	if err != nil {
		t.Error(err)
	}

	applications, err := data.GetApplications(db, user_id)
	if err != nil {
		t.Error(err)
	}

	for index, app := range applications {

		// get the application data
		applicationData, err := data.GetApplicationData(db, user_id, strconv.Itoa(app.ID))
		if err != nil {
			t.Error(err)
		}

		if len(applicationData) <= 0 {
			t.Error("No application data found")
		}

		// check application name
		decryptedUsername, err := tools.Decrypt(applicationData[0].Username, key)
		if err != nil {
			t.Error(err)
		}
		decryptedPassword, err := tools.Decrypt(applicationData[0].Password, key)
		if err != nil {
			t.Error(err)
		}

		if decryptedUsername != application_username {
			t.Error("Expected username: ", application_username, " got ", applicationData[index].Username)
		}

		// check application password
		if decryptedPassword != application_password {
			t.Error("Expected password: ", application_password, " got ", applicationData[index].Password)
		}
	}

	//close the database connection
	db.Close()
}

func TestDatabaseExistence(t *testing.T) {
	if data.IsDatabaseExit(database) {
		t.Log("Database exists")
		cmd := exec.Command("rm", database)
		err := cmd.Run()
		if err != nil {
			t.Error("Error removing database: ", err)
		}
	} else {
		t.Error("Database does not exist")
	}
}
