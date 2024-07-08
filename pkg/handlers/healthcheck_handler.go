package handlers

import (
	"net/http"
)

type Response struct {
    Message string `json:"message"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

    response := Response{
        Message: "Hello!",
    }

	Respond(w, response, 200)
}
