package main

import (
	"go-start/pkg/db"
	"go-start/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8000"

	db.InitDB()

    mux := &http.ServeMux{}

	mux.HandleFunc("/", handlers.HealthCheckHandler)
	mux.HandleFunc("/users", handlers.GetUsersHandler)
	mux.HandleFunc("/users/{id}", handlers.GetUserHandler)
	mux.HandleFunc("/users/create", handlers.CreateUserHandler)
	mux.HandleFunc("/users/delete/{id}", handlers.DeleteUserHandler)

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
