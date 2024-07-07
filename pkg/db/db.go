package db

import (
	"database/sql"
	"errors"
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

func GetUsers() ([]models.User, error) {
    rows, err := DB.Query("SELECT * from user")
    if err != nil {
        return nil, errors.New("Could not prepare statement.")
    }

    defer rows.Close()

    var users []models.User

    for rows.Next() {
        var user models.User
        err = rows.Scan(&user.ID, &user.Name)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
        return nil, errors.New("An error occured.")
    }

    return users, nil
}
