package main

import (
	"net/http"

	"github.com/douttnus/restapi/handlers"
)

func main() {
	// users
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/go", handlers.GetUser)

	http.HandleFunc("/users/new", handlers.CreateUser)
	http.HandleFunc("/users/update", handlers.UpdateUser)
	http.HandleFunc("/users/delete", handlers.DeleteUser)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
