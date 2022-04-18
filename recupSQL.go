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
		var avatar string
		typee := "guest"
		err = rows.Scan(&id, &uid, &username, &email, &passwd, &avatar, &typee)
		Debug(err)
		newTab = append(newTab, User{Id: id, Uid: uid, Email: email, Username: username, Passwd: passwd, Avatar: avatar, Type: typee})
	}

	return newTab
}

func RecupPost() []Post {

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)

	rows, err := db.Query("SELECT * FROM posts")
	Debug(err)

	err = rows.Err()
	Debug(err)

	var newTab []Post
	for rows.Next() {
		var pid int
		var content string
		var category string
		var uid string

		var counter int

		err = rows.Scan(&pid, &content, &category, &uid)
		Debug(err)
		count, err := db.Query(`SELECT COUNT(*) FROM likes WHERE pid = ?`, pid)
		Debug(err)
		for count.Next() {
			_ = count.Scan(&counter)
		}
		newTab = append(newTab, Post{Pid: pid, Content: content, Category: category, Uid: uid, LikeActive: "unlike", Like: counter})
	}

	for post := range newTab {
		for _, user := range Users {
			if newTab[post].Uid == user.Uid {
				newTab[post].User = user
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

	for _, comment := range newTab {
		for _, user := range Users {
			if comment.Uid == user.Uid {
				comment.User = user
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

func RecupLike(Uid string)  {
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM likes"
	rows, err := db.Query(selection)
	Debug(err)
	err = rows.Err()
	Debug(err)
	for rows.Next() {
		var id int
		var uid string
		var pid int
		err = rows.Scan(&id, &uid, &pid)
		Debug(err)

		for post := range Posts {
			if pid == Posts[post].Pid && uid == Uid {
				Posts[post].LikeActive = "like"
			}
		}
	}
}
