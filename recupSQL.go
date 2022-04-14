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
		var id int
		var uid string
		var username string
		var email string
		var passwd string
		err = rows.Scan(&id, &uid, &username, &email, &passwd)
		Debug(err)
		newTab = append(newTab, User{Id: id, Uid: uid, Email: email, Username: username, Passwd: passwd})
	}
	err = rows.Err()
	Debug(err)

	return newTab
}

func RecupPost() []Post {
	var newTab []Post
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM posts"
	rows, err := db.Query(selection)
	Debug(err)
	err = rows.Err()
	Debug(err)
	for rows.Next() {

		//posts
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
	var newTab []Comment
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM comments"
	rows, err := db.Query(selection)
	Debug(err)
	err = rows.Err()
	Debug(err)
	for rows.Next() {

		//comments
		var cid int
		var pid int
		var comment string

		err = rows.Scan(&cid, &pid, &comment)
		Debug(err)
		newTab = append(newTab, Comment{Cid: cid, Pid: pid, Content: comment})
		// notfind := true
		// fmt.Println(newTab)

		// // for i := range newTab {
		// // 	if newTab[i].Pid == pid {
		// // 		notfind = false
		// // 		newTab[i].Comments = append(newTab[i].Comments, comment)
		// // 	}
		// // }
		// // if notfind {
		// // 	newTab = append(newTab, Post{Pid: pid, Content: content, Comments: []string{comment}})
		// // }

	}

	err = rows.Err()
	Debug(err)

	return newTab
}
