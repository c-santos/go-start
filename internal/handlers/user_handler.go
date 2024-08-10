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
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, users, 200)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.PathValue("id")

	user, err := db.GetUser(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, user, 200)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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
	user_id := r.PathValue("id")

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	user, err := db.UpdateUser(user_id, updatedUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	Respond(w, user, 200)
}

func CreateUserNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from path params.
	user_id := r.PathValue("user_id")
	int_user_id, err := strconv.Atoi(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode response body to note variable.
	var note models.Note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign user_id.
	*&note.UserID = int_user_id

	created_note, err := db.CreateNote(note)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	Respond(w, created_note, 201)
}

func GetUserNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from path params.
	user_id := r.PathValue("user_id")
	int_user_id, err := strconv.Atoi(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notes, err := db.GetNotesByUserId(int_user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, notes, 200)
}

func DeleteUserNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from path params.
	user_id := r.PathValue("user_id")
	note_id := r.PathValue("note_id")

	note, err := db.GetNote(note_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteNote(user_id, note_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, note, 200)
}

func UpdateUserNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from path params.
	user_id := r.PathValue("user_id")
	note_id := r.PathValue("note_id")

	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)

	updated_note, err := db.UpdateNote(note_id, user_id, note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, updated_note, 200)
}
