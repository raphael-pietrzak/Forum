package forum

import (
	"database/sql"
	"fmt"

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
	db.Close()

	return newTab
}

func RecupPost() []Post {
	var newTab []Post
	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	selection := "SELECT * FROM posts JOIN comments ON posts.pid = comments.pid"
	rows, err := db.Query(selection)
	Debug(err)
	err = rows.Err()
	Debug(err)
	fmt.Print("test : ")
	for rows.Next() {

		//posts
		var pid int
		var uid string
		var content string

		//comments
		var cid int
		var pid2 int
		var uid2 string
		var comment string

		err = rows.Scan(&pid, &uid, &content, &cid, &pid2, &uid2, &comment)
		notfind := true

		fmt.Println(pid, content, cid, comment, pid2, uid, uid2)

		for i := range newTab {
			if newTab[i].Pid == pid {
				notfind = false
				newTab[i].Comments = append(newTab[i].Comments, comment)
			}
		}
		if notfind {
			newTab = append(newTab, Post{Pid: pid, Content: content, Comments: []string{comment}})
		}

		Debug(err)
	}

	fmt.Println(newTab)
	err = rows.Err()
	Debug(err)
	db.Close()

	return newTab
}
