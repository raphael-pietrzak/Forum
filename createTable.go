package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func TableCreation() {
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user 
	(
		id INTEGER PRIMARY KEY, 
		uid TEXT, 
		username TEXT, 
		email TEXT, 
		passwd TEXT,
		avatar TEXT,
		type TEXT

	);`)

	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts 
	(
		pid INTEGER PRIMARY KEY, 
		content TEXT,
		category TEXT,
		contentPhoto TEXT,
		uid TEXT

	);`)

	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments 
	(
		cid INTEGER PRIMARY KEY,
		pid INTEGER,
		content TEXT,
		uid TEXT
	); `)

	Debug(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS likes 
	(
		id INTEGER PRIMARY KEY, 
		uid TEXT, 
		pid INTEGER,
		cid INTEGER,
		notification BOOL
	);`)

	Debug(err)
}
