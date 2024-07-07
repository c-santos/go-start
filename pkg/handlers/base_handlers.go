package handlers

import (
	"encoding/json"
	"net/http"
    "go-start/pkg/models"
)

func Respond(w http.ResponseWriter, response models.Response) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Message: "Hello, you've called me!",
		Status:  200,
	}

	Respond(w, response)
}
