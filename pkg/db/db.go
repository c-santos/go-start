package db

import (
	"database/sql"
	"go-start/pkg/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS user (
        id INTEGER NOT NULL PRIMARY KEY,
        name TEXT
    );
`
	_, err := DB.Exec(sqlStmt)

	if err != nil {
		log.Fatalf("%qL %s\n", err, sqlStmt)
		return
	} else {
		log.Println("Database created successfully.")
	}
}

func CreateUser(user models.User) error {
    stmt, err := DB.Prepare("INSERT INTO user(name) VALUES(?)")
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(user.Name)
    if err != nil {
        return err
    }

    return nil
}
