package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Database transaction
//-----------------------

// 1. User creates account
// 2. Create a wallet for the user
// 3. Want to top up the wallet for the user
// 4. You want to write a transaction log

var userSchema = `
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
	Profile        Profile   `json:"profile"`
}

var profilSchema = `
CREATE TABLE IF NOT EXISTS profile (
    user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    avatar TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type Profile struct {
	UserID  int       `json:"user_id"`
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
}

func main() {
	dbName := "users_database.db"
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createUserTable(db)
	fmt.Println("users table is created")
	createProfileTable(db)
	fmt.Println("profile table is created")
	_, err = createUser(db, "test with defer", "test@tets.s", "qkjsq", "kjqwkq")
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTable(db *sql.DB) {
	_, err := db.Exec(userSchema)
	if err != nil {
		log.Fatal(err)
	}

}

func createProfileTable(db *sql.DB) {
	_, err := db.Exec(profilSchema)
	if err != nil {
		log.Fatal(err)
	}

}

// Begin , Rollback or Commit

func createUser(db *sql.DB, name string, email string, password string, avatar string) (int64, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	userStmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer userStmt.Close()
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := userStmt.Exec(name, email, string(hp))
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	profileStm, err := tx.PrepareContext(ctx, `INSERT INTO profile (user_id, avatar) VALUES( ?, ?)`)
	if err != nil {
		return 0, err
	}

	defer profileStm.Close()
	_, err = profileStm.Exec(userID, avatar)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return userID, nil
}
