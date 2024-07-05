package main

import (
	"log"
	"net/http"
    "go-start/pkg/handlers"
    "go-start/pkg/db"
)

func main() {
    port := ":8000"

    db.InitDB()

	http.HandleFunc("/", handlers.HealthCheckHandler)
	http.HandleFunc("/users", handlers.UserHandler)
    http.HandleFunc("/users/create", handlers.CreateUserHandler)

    log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
