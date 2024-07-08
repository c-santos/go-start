package handlers

import (
	"net/http"
)


func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    type Response struct {
        Message string `json:"message"`
    }

    response := Response{
        Message: "Hello!",
    }

	Respond(w, response)
}
