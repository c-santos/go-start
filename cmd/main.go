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
	// http.HandleFunc("/hello/{name}", helloHandler)

    log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
