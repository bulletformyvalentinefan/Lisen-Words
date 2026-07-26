package db

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, username, password string) error {

	q := `INSERT INTO users (username, password) VALUES (?, ?)`

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(q, username, string(passwordHash))
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(db *sql.DB, username, password string) (int64, error) {

	q := `SELECT id, password FROM users WHERE username = ?`

	var id int64
	var storeHash string

	err := db.QueryRow(q, username).Scan(&id, &storeHash)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storeHash), []byte(password))
	if err != nil {
		return 0, err 
	}

	return id, nil
}
