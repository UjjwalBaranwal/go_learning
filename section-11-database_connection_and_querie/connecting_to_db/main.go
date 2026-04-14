package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password BLOB NOT NULL, -- Storing as BLOB for byte slice
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	dbName := "data.db"
	// dont do this is prod environment
	_ = os.Remove(dbName)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("closing database connection")
		if err := db.Close(); err != nil {
			log.Printf("error closingg the database connection : %v\n", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection established")
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table is Created")
}
