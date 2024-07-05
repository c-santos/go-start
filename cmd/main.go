package main

import (
	"log"
	"net/http"
    "go-start/pkg/handlers"
)

func main() {
    port := ":8000"

	http.HandleFunc("/", handlers.HealthCheckHandler)
	// http.HandleFunc("/hello/{name}", helloHandler)

    log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
