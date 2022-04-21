package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func RecupUser() {
	//users
	var id int
	var uid string
	var username string
	var email string
	var passwd string
	var avatar string
	var typee string

	typee = "guest"

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)

	rows, err := db.Query("SELECT * FROM user")
	Debug(err)

	err = rows.Err()
	Debug(err)

	var newTab []User

	for rows.Next() {
		err = rows.Scan(&id, &uid, &username, &email, &passwd, &avatar, &typee)
		Debug(err)
		newTab = append(newTab, User{Id: id, Uid: uid, Email: email, Username: username, Passwd: passwd, Avatar: avatar, Type: typee})
	}

	Users = newTab
}

func RecupPost() {

	//Posts
	var pid int
	var content string
	var category string
	var uid string
	var counter int
	var contentPhoto string

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)

	rows, err := db.Query("SELECT * FROM posts")
	Debug(err)

	err = rows.Err()
	Debug(err)

	var newTab []Post
	for rows.Next() {

		err = rows.Scan(&pid, &content, &category, &contentPhoto, &uid)
		Debug(err)
		count, err := db.Query(`SELECT COUNT(*) FROM likes WHERE pid = ?`, pid)
		Debug(err)
		for count.Next() {
			_ = count.Scan(&counter)
		}
		newTab = append(newTab, Post{Pid: pid, Content: content, Category: category, Uid: uid, LikeActive: "unlike", Like: counter, ContentPhoto: contentPhoto})
	}

	for post := range newTab {
		for _, user := range Users {
			if newTab[post].Uid == user.Uid {
				newTab[post].User = user
			}
		}
	}

	Posts = newTab
}

func RecupComment() {
	//comments
	var cid int
	var pid int
	var comment string
	var uid string
	var user_comment User
	var counter int

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	rows, err := db.Query("SELECT * FROM comments")
	Debug(err)
	err = rows.Err()
	Debug(err)

	for rows.Next() {

		err = rows.Scan(&cid, &pid, &comment, &uid)
		Debug(err)

		count, err := db.Query(`SELECT COUNT(*) FROM likes WHERE cid = ?`, cid)
		Debug(err)
		for count.Next() {
			_ = count.Scan(&counter)
		}

		for _, user := range Users {
			if uid == user.Uid {
				user_comment = user
				break
			}
		}
		for post := range Posts {
			if pid == Posts[post].Pid {
				Posts[post].Comments = append(Posts[post].Comments, Comment{Uid: uid, Pid: pid, Content: comment, Cid: cid, User: user_comment, LikeActive: "unlike", Like: counter})
			}
		}
	}
}

func RecupLike(Uid string) {
	//like
	var id int
	var uid string
	var pid int
	var cid int

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	rows, err := db.Query("SELECT * FROM likes")
	Debug(err)
	err = rows.Err()
	Debug(err)

	for rows.Next() {

		err = rows.Scan(&id, &uid, &pid, &cid)
		Debug(err)
		if uid == Uid {
			if cid == 0 {
				Posts[pid-1].LikeActive = "like"
			} else {
				for i, post := range Posts {
					for j := range post.Comments {
						if Posts[i].Comments[j].Cid == cid {
							Posts[i].Comments[j].LikeActive = "like"
						}
					}
				}
			}
		}
	}
}
