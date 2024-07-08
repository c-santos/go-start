package handlers

import (
	"encoding/json"
	"go-start/pkg/db"
	"go-start/pkg/models"
	"net/http"
	"strconv"
)

type CreateUserDto struct {
	Name string `json:"name"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, users, 200)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user_id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Could not convert id to int", http.StatusInternalServerError)
        return
	}

	user, err := db.GetUser(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	Respond(w, user, 200)
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

	Respond(w, user, 201)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete { 
        http.Error(w, "Method not allowed", http.StatusBadRequest) 
        return
    }

    user_id := r.PathValue("id")

    err := db.DeleteUser(user_id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    message := Response{
        Message: "Deleted user.",
    }

    Respond(w, message, 200)
}

