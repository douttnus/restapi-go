package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/douttnus/restapi/models"
)

// Data created for testing is fictitious.
// (Whether you want to enter your own details here or leave it blank is your choice).
var users = []models.User{
	{ID: 1, Name: "Cammily Burges", Email: "camillyburges@gmail.com", Password: "g57K_97gp_^#2VZcr"},
	{ID: 2, Name: "Douttnus Mile", Email: "douxnuus@tuta.com", Password: "L9gk<9q7s%0a,5"},
}

// METHOD GET
// We are getting all the users that should be being delivered from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user) // let's read the body of the request and assimilate it with the user's struct
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uExist := false
	for _, u := range users {
		if u.ID == user.ID {
			user = u
			uExist = true
			break
		}
	}

	if !uExist {
		http.Error(w, "User with ID not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK) // If everything goes well, we will return a status code of 201
	json.NewEncoder(w).Encode(user)
}

// METHOD POST
// Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User // We call the users struct

	err := json.NewDecoder(r.Body).Decode(&newUser) // let's read the body of the request and assimilate it with the user's struct
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated) // If everything goes well, we will return a status code of 201
	json.NewEncoder(w).Encode(newUser)
}

// METHOD PUT
// Update a user's data if it exists (by ID)
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser models.User

	// transforms data into JSON
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// checks if the ID passed in the request body exists somewhere in the slice
	uExist := false
	for _, u := range users {
		if u.ID == updateUser.ID {
			uExist = true
			break
		}
	}

	if !uExist {
		http.Error(w, "User with ID not found", http.StatusBadRequest)
	}

	// runs through the entire slice and assigns new data to the existing slice
	for i, user := range users {
		if user.ID == updateUser.ID {
			users[i] = updateUser
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateUser)
}

// METHOD DELETE
// Deletes a specific user by passing the ID and password in the body of the request
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser models.User

	err := json.NewDecoder(r.Body).Decode(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uExist := false
	for _, u := range users {
		if u.ID == deleteUser.ID {
			uExist = true
			break
		}
	}

	if !uExist {
		http.Error(w, "User with ID not found", http.StatusBadRequest)
	}

	// runs through the entire slice and assigns new data to the existing slice
	upUser := []models.User{}
	for _, user := range users {
		if user.ID != deleteUser.ID && user.Password == deleteUser.Password {
			upUser = append(upUser, user)
		}
	}
	users = upUser

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteUser)
}
