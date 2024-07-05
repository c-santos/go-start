package handlers

import (
	"encoding/json"
	"go-start/pkg/db"
	"go-start/pkg/models"
	"net/http"
)

type CreateUserDto struct {
    Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    users := []models.User{
        {ID: 1, Name: "John"},
        {ID: 2, Name: "Bob"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusBadRequest)
        return
    }

    var user models.User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = db.CreateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
