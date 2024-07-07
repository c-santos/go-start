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

	http.HandleFunc("/", handlers.HealthCheckHandler)
	http.HandleFunc("/users", handlers.GetUsersHandler)
	http.HandleFunc("/users/create", handlers.CreateUserHandler)

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
