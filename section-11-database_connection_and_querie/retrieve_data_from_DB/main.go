package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

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
	// createTable(db)
	// fmt.Println("Table is Created")
	// id, err := createUser(db, "Ujjwal", "u@gmail.com", "i_am_anti_hero")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("id : %d\n", id)
	// }

	user, err := getUserByEmail(db, "u@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := json.MarshalIndent(user, "", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
	users, err := getAllUser(db)
	if err != nil {
		log.Fatal(err)
	}

	bs, err = json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
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

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `Select id,name,email,hashed_password,created_at FROM users where email = ?`
	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getAllUser(db *sql.DB) ([]User, error) {
	stmt := `SELECT id,name,email,hashed_password,created_at from users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
