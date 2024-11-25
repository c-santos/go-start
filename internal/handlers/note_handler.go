package handlers

import (
	"encoding/json"
	"go-start/internal/db"
	"go-start/internal/models"
	"net/http"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := db.GetNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, notes, 200)
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err = db.CreateNote(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, note, 201)
}
