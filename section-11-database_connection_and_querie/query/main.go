package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL, 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	dbName := "users_database.db"
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection established")
	createTable(db)
	fmt.Println("Table is Created")
	id, err := createUser(db, "Ujjwal", "u@gmail.com", "i_am_anti_hero")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("id : %d\n", id)
	}
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

}

func createUser(db *sql.DB, name string, email string, password string) (int64, error) {
	stmt := `INSERT INTO users (name,email,hashed_password) VALUES (?,?,?)`
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	result, err := db.Exec(stmt, name, email, hp)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return result.LastInsertId()
}
