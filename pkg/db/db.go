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

func CreateUser(user models.User) (models.User, error) {
	stmt, err := DB.Prepare("INSERT INTO user(name) VALUES(?)")
	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name)
	if err != nil {
		return models.User{}, err
	}

	var user_id int64
	user_id, err = result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	rows := DB.QueryRow("SELECT * FROM user WHERE id = ?", user_id)

	var created_user models.User
	err = rows.Scan(&created_user.ID, &created_user.Name)
	if err != nil {
		return models.User{}, err
	}

	return created_user, nil
}

func GetUsers() (models.Users, error) {
	rows, err := DB.Query("SELECT * from user")
	if err != nil {
		return nil, errors.New("Query error.")
	}

	var users models.Users
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(user_id string) (models.User, error) {
	row := DB.QueryRow("SELECT * from user where id = ?", user_id)

	if row.Err() != nil {
		return models.User{}, errors.New("Query unsuccessful.")
	}

	var user models.User

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return models.User{}, errors.New("Row scan failed.")
	}

	return user, nil
}

func DeleteUser(user_id string) error {
	stmt, err := DB.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user_id string, user models.User) (models.User, error) {
	stmt, err := DB.Prepare("UPDATE user SET name = ? WHERE id = ?")
	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user_id)
	if err != nil {
		return models.User{}, err
	}

	rows := DB.QueryRow("SELECT * from user WHERE id = ?", user_id)

	var updatedUser models.User
	err = rows.Scan(&updatedUser.ID, &updatedUser.Name)

	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}
