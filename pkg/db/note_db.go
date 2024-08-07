package db

import (
	"errors"
	"go-start/pkg/models"
	"log"
)

func CreateNote(note models.Note) (models.Note, error) {
	stmt, err := DB.Prepare("INSERT INTO note(title, body, user_id) VALUES(?, ?, ?)")
	if err != nil {
		return models.Note{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(note.Title, note.Body, note.UserID)
	if err != nil {
		return models.Note{}, err
	}

	var note_id int64
	note_id, err = result.LastInsertId()
	if err != nil {
		return models.Note{}, err
	}

	rows := DB.QueryRow("SELECT * FROM note WHERE id = ?", note_id)

	var created_note models.Note
	err = rows.Scan(&created_note.ID, &created_note.Title, &created_note.Body, &created_note.UserID)
	if err != nil {
		return models.Note{}, err
	}

	return created_note, nil
}

func GetNotes() (models.Notes, error) {
	rows, err := DB.Query("SELECT * from note")
	if err != nil {
		return nil, errors.New("Query error.")
	}

	var notes models.Notes
	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.ID, &note.Title, &note.Body, &note.UserID)
		if err != nil {
			log.Printf("[db.GetNotes] %s", err)
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func GetNote(note_id string) (models.Note, error) {
	row := DB.QueryRow("SELECT * from note where id = ?", note_id)
	if row.Err() != nil {
		return models.Note{}, errors.New("Query unsuccessful.")
	}

	var note models.Note

	err := row.Scan(&note.ID, &note.Title, &note.Body, &note.UserID)
	if err != nil {
		return models.Note{}, errors.New("Row scan failed.")
	}

	return note, nil
}

func DeleteNote(user_id string, note_id string) error {
	stmt, err := DB.Prepare("DELETE FROM note WHERE id = ? AND user_id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user_id, note_id)
	if err != nil {
		return err
	}

	return nil
}

func GetNotesByUserId(user_id int) (models.Notes, error) {
	rows, err := DB.Query("SELECT * from note WHERE user_id = ?", user_id)
	if err != nil {
		return nil, errors.New("Query error.")
	}

	var notes models.Notes
	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.ID, &note.Title, &note.Body, &note.UserID)
		if err != nil {
			log.Printf("[db.GetNotes] %s", err)
		}
		notes = append(notes, note)
	}

	return notes, nil
}
