package data

type User struct {
	ID       int
	Username string
	Password string
}

type Application struct {
	ID   int
	Name string
	URL  string
}

type ApplicationData struct {
	UserID        int
	ApplicationID int
	Username      []byte
	Password      []byte
	DataCreated   string
	LastUpdated   string
}
