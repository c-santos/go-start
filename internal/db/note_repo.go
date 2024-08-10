package db

import (
	"errors"
	"go-start/internal/models"
	"log"
)

func InitNote() string {
	stmt := `
    CREATE TABLE IF NOT EXISTS note (
        id INTEGER NOT NULL PRIMARY KEY,
        title TEXT,
        body TEXT,
        user_id INTEGER NOT NULL,
        FOREIGN KEY(user_id) REFERENCES user(id)
    );`

	return stmt
}

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
	row := DB.QueryRow("SELECT * from note WHERE id = ?", note_id)
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

func UpdateNote(note_id string, user_id string, note models.Note) (models.Note, error) {
	var query string
	var args []interface{}
	log.Println(note)

	if note.Body == "" && note.Title == "" {
		return models.Note{}, errors.New("Nothing to update.")
	}

	if note.Title != "" {
		query = "UPDATE note SET title = ?"
		args = append(args, note.Title)
	}

	if note.Body != "" {
		if query != "" {
			query += ", body = ?"
		} else {
			query = "UPDATE note SET body = ?"
		}
		args = append(args, note.Body)
	}

	query += " WHERE id = ? and user_id = ?"
	args = append(args, note_id, user_id)

	log.Println(query)

	stmt, err := DB.Prepare(query)
	if err != nil {
		return models.Note{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return models.Note{}, err
	}

	rows := DB.QueryRow("SELECT * from note WHERE id = ?", note_id)

	var updated_note models.Note
	err = rows.Scan(&updated_note.ID, &updated_note.Title, &updated_note.Body, &updated_note.UserID)
	if err != nil {
		return models.Note{}, err
	}

	return updated_note, nil
}
