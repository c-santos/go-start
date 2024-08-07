package main

import (
	"go-start/pkg/db"
	"go-start/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8000"

	db.InitDB()

	mux := &http.ServeMux{}

	mux.HandleFunc("GET /health", handlers.HealthCheckHandler)
	mux.HandleFunc("GET /users", handlers.GetUsersHandler)
	mux.HandleFunc("GET /users/{id}", handlers.GetUserHandler)
	mux.HandleFunc("POST /users", handlers.CreateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUserHandler)
	mux.HandleFunc("PATCH /users/{id}", handlers.UpdateUserHandler)

	mux.HandleFunc("GET /users/{user_id}/notes", handlers.GetUserNotesHandler)
	mux.HandleFunc("POST /users/{user_id}/notes", handlers.CreateUserNoteHandler)
	mux.HandleFunc("DELETE /users/{user_id}/notes/{note_id}", handlers.DeleteUserNoteHandler)

	mux.HandleFunc("GET /notes", handlers.GetNotesHandler)
	mux.HandleFunc("POST /notes", handlers.CreateNoteHandler)

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
