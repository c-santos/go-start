package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

    helloHandler := func (w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello, you've called me!")
    }

    userHandler := func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s\n", r.PathValue("name"))
    }

    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/hello/{name}", userHandler)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
