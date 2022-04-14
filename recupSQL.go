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
		var category string
		var uid string

		err = rows.Scan(&pid, &content, &category, &uid)
		Debug(err)
		newTab = append(newTab, Post{Pid: pid, Content: content, Category: category, Uid: uid})
	}

	err = rows.Err()
	Debug(err)

	for post := range newTab {
		for _, user := range Users {
			if newTab[post].Uid == user.Uid {
				newTab[post].Username = user.Username
			}
		}
	}

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
		var uid string

		err = rows.Scan(&cid, &pid, &comment, &uid)
		Debug(err)
		newTab = append(newTab, Comment{Cid: cid, Pid: pid, Content: comment, Uid: uid})

		
	}

	err = rows.Err()
	Debug(err)
	for _,comment := range newTab {
		for _,user := range Users {
			if comment.Uid == user.Uid {
				comment.Username = user.Username
			}
		}
		for post := range Posts {
			if comment.Pid == Posts[post].Pid {
				Posts[post].Comments = append(Posts[post].Comments, comment)
			}
		}
	}
	return newTab
}
