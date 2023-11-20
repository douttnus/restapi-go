package models

// Creating the necessary structures for our RESTAPI
// This will be the data we will use (which would come from a database)
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
