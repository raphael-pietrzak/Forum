package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func RecupUser() []User {
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM user"
	rows, err := db.Query(selection)
	Debug(err)

	err = rows.Err()
	Debug(err)
	var newTab []User
	for rows.Next() {
		var uid int
		var username string
		var email string
		var passwd string
		err = rows.Scan(&uid, &username, &email, &passwd)
		Debug(err)
		newTab = append(newTab, User{Uid: uid, Email: email, Username: username, Passwd: passwd})
	}
	err = rows.Err()
	Debug(err)
	return newTab
}

func RecupPost() []Post {
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM posts"
	rows, err := db.Query(selection)
	Debug(err)

	err = rows.Err()
	Debug(err)
	var newTab []Post
	for rows.Next() {
		var pid int
		var content string
		err = rows.Scan(&pid, &content)
		Debug(err)
		newTab = append(newTab, Post{Pid: pid, Content: content})
	}
	err = rows.Err()
	Debug(err)
	return newTab
}

func RecupComment() []Comment {
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM comment"
	rows, err := db.Query(selection)
	Debug(err)

	err = rows.Err()
	Debug(err)
	var newTab []Comment
	for rows.Next() {
		var cid int
		var content string
		err = rows.Scan(&cid, &content)
		Debug(err)
		newTab = append(newTab, Comment{Cid: cid, Content: content})
	}
	err = rows.Err()
	Debug(err)
	return newTab
}
