package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, user *User) error {
	query := `
        INSERT INTO users 
        (username, password_hash) 
        VALUES (?, ?)`

	_, err := db.Exec(query,
		user.Username,
		user.PasswordHash,
	)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (User, error) {
	query := "SELECT id, username, password_hash FROM users WHERE username = ?"
	var user User
	err := db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
	)
	return user, err
}
