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

	user_id := r.PathValue("id")

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

	user, err = db.CreateUser(user)
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

    user, err := db.GetUser(user_id)

	err = db.DeleteUser(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, user, 200)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
        http.Error(w, "Method not allowed", http.StatusBadRequest)
        return
    }

    user_id := r.PathValue("id")

    var updatedUser models.User
    json.NewDecoder(r.Body).Decode(&updatedUser)

    user, err := db.UpdateUser(user_id, updatedUser)
    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    Respond(w, user, 200)
}

