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
    table_init_stmt := InitUser() + InitNote()

	_, err := DB.Exec(table_init_stmt)

	if err != nil {
		log.Fatalf("%qL %s\n", err, table_init_stmt)
		return
	} else {
		log.Println("Database created successfully.")
	}
}
