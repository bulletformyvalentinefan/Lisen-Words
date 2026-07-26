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
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tracks (
			id INTEGER PRIMARY KEY, 
			title TEXT NOT NULL,
			duration INTEGER NOT NULL,
			preview TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS favorites (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			track_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(track_id) REFERENCES tracks(id)
		);

		CREATE TABLE IF NOT EXISTS playlists (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS playlist_tracks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			playlist_id INTEGER NOT NULL,
			track_id INTEGER NOT NULL,
			FOREIGN KEY(playlist_id) REFERENCES playlists(id),
			FOREIGN KEY(track_id) REFERENCES tracks(id)
		);`

    if _, err := db.Exec(q); err != nil {
        return nil, err
    }
	
    return db, nil
}