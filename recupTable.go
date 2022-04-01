package forum

import (
	"database/sql"
	_ "log"

	_ "github.com/mattn/go-sqlite3"
)


func RecupUser() []User {
	db, err := sql.Open("sqlite3", "./database.db")
	selection := "SELECT * FROM user"
	rows, err := db.Query(selection)
	debug(err)

	err = rows.Err()
	debug(err)
	var newTab []User
	for rows.Next() {
		var uid int
		var username string
		var email string
		var passwd string
		err = rows.Scan(&uid, &username, &email, &passwd)
		debug(err)
		newTab = append(newTab, User{Uid: uid, Email: email, Username: username, Passwd: passwd})
	}
	err = rows.Err()
	debug(err)
	return newTab
}

func RecupPost() []Post {
	db, err := sql.Open("sqlite3", "./database.db")
	selection := "SELECT * FROM posts"
	rows, err := db.Query(selection)
	debug(err)

	err = rows.Err()
	debug(err)
	var newTab []Post
	for rows.Next() {
		var pid int
		var content string
		err = rows.Scan(&pid, &content)
		debug(err)
		newTab = append(newTab, Post{Pid: pid, Content: content})
	}
	err = rows.Err()
	debug(err)
	return newTab
}

func RecupComment() []Comment {
	db, err := sql.Open("sqlite3", "./database.db")
	selection := "SELECT * FROM comment"
	rows, err := db.Query(selection)
	debug(err)

	err = rows.Err()
	debug(err)
	var newTab []Comment
	for rows.Next() {
		var cid int
		var content string
		err = rows.Scan(&cid, &content)
		debug(err)
		newTab = append(newTab, Comment{Cid: cid, Content: content})
	}
	err = rows.Err()
	debug(err)
	return newTab
}
