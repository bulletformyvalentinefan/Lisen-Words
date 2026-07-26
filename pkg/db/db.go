package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func InitDB(dbPath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite", dbPath)
    if err != nil {
        return nil, err
    }

    q := `
    -- Tus CREATE TABLE IF NOT EXISTS aquí
    `

    if _, err := db.Exec(q); err != nil {
        return nil, err
    }

    return db, nil
}