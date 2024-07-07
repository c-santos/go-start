package handlers

import (
	"encoding/json"
	"go-start/pkg/models"
	"net/http"
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
