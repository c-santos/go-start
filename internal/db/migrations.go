package db

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

func InitMigrations() string {
	stmt := `
	CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT NOT NULL
	)
	`

	return stmt
}

func Migrate() {
	// set migrations directory
	user, err := user.Current()
	if err != nil {
		log.Fatalf("%s ", err)
	}

	base := user.HomeDir
	fmt.Println(base)

	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	migrationsDir := path.Join(ex, "./migrations")
	fmt.Println(migrationsDir)

	// fs.WalkDir(migrationsDir, ".", func(path string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	fmt.Println(path)
	// 	return nil
	// })
}

