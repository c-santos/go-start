package handlers

import (
	"net/http"
    "go-start/pkg/models"
)


func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Message: "Hello, you've called me!",
		Status:  200,
	}

	Respond(w, response)
}
