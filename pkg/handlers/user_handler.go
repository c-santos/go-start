package handlers

import (
	"encoding/json"
	"go-start/pkg/models"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    users := []models.User{
        {ID: 1, Name: "John"},
        {ID: 2, Name: "Bob"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
