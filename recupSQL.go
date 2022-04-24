package forum

import (
	"database/sql"
	"fmt"

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
	typee := "guest"

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
	var date string
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

		err = rows.Scan(&pid, &content, &category, &date, &contentPhoto, &uid)
		Debug(err)
		date = TimeSince(date)

		count, err := db.Query(`SELECT COUNT(*) FROM likes WHERE pid = ?`, pid)
		Debug(err)
		for count.Next() {
			_ = count.Scan(&counter)
		}
		// сounter = 1465
		finalcounter := Adaptlikes(counter)
		newTab = append(newTab, Post{Pid: pid, Content: content, Category: category, Uid: uid, LikeActive: "unlike", Like: finalcounter, Date: date, ContentPhoto: contentPhoto, Like2: counter})
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
	var date string
	var counter int

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	rows, err := db.Query("SELECT * FROM comments")
	Debug(err)
	err = rows.Err()
	Debug(err)

	for rows.Next() {

		err = rows.Scan(&cid, &pid, &comment, &date, &uid)
		Debug(err)

		date = TimeSince(date)
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
		// counter = 3987888

		finalcounter := Adaptlikes(counter)
		for post := range Posts {
			if pid == Posts[post].Pid {
				Posts[post].Comments = append(Posts[post].Comments, Comment{Uid: uid, Pid: pid, Content: comment, Cid: cid, User: user_comment, LikeActive: "unlike", Like: finalcounter, Date: date})
			}
		}
	}
}

func RecupLike(User User) {
	//like
	var id int
	var uid string
	var user_liked string
	var vu bool
	var pid int
	var cid int
	// var uidreceveur string

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)
	rows, err := db.Query("SELECT * FROM likes")
	Debug(err)
	err = rows.Err()
	Debug(err)
	Data.Notifications = []string{}
	for rows.Next() {

		err = rows.Scan(&id, &uid, &user_liked, &vu, &pid, &cid)
		Debug(err)
		if !vu && User.Uid == user_liked {
			for _,user := range Users {
				if user.Uid == uid {
					Data.Notifications = append(Data.Notifications,user.Username + " liked your post")
					break
				}
			}
		}
		if uid == User.Uid {
			if cid == 0 {
				for post := range Posts {
					if pid == Posts[post].Pid {
						Posts[post].LikeActive = "like"
					}
				}
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
	Data.Notifs = len(Data.Notifications)

}

func Adaptlikes(like int) string {
	var finalcounter string
	if like > 1000000000 {
		finalcounter = fmt.Sprintf("%.2f", float64(like)/1000000000) + "Md"
	} else if like > 1000000 {
		finalcounter = fmt.Sprintf("%.2f", float64(like)/1000000) + "M"
	} else if like > 1000 {
		finalcounter = fmt.Sprintf("%.2f", float64(like)/1000) + "k"
	} else {
		finalcounter = fmt.Sprintf("%d", like)
	}
	return finalcounter
}
