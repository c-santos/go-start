package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
    CREATE TABLE IF NOT EXISTS note (
        id INTEGER NOT NULL PRIMARY KEY,
        title TEXT,
        body TEXT,
        user_id INTEGER NOT NULL,
        FOREIGN KEY(user_id) REFERENCES user(id)
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
