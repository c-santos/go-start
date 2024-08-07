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

	mux.HandleFunc("GET /health", handlers.HealthCheckHandler)
	mux.HandleFunc("GET /users", handlers.GetUsersHandler)
	mux.HandleFunc("GET /users/{id}", handlers.GetUserHandler)
	mux.HandleFunc("POST /users", handlers.CreateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUserHandler)
	mux.HandleFunc("PATCH /users/{id}", handlers.UpdateUserHandler)

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
