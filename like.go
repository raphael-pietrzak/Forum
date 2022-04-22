package forum

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

func LikePosts(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)
	var pid int
	var cid int
	var notification bool

	likerecup := strings.Split(r.Form.Get("like"), " ")
	pid_cid, err := strconv.Atoi(likerecup[0][1:])
	Debug(err)
	which_pid_or_cid, err := strconv.Atoi(likerecup[1])
	Debug(err)
	like := likerecup[5]

	if which_pid_or_cid == 0 {
		pid = pid_cid
	} else {
		cid = pid_cid
	}

	if like == "like" {
		notification = true
		_, err := db.Exec(`INSERT OR IGNORE INTO likes ('pid', 'uid', 'cid','notification') VALUES (?, ?, ?, ?)`, pid, GetUserByCookies(w, r).Uid, cid, notification)
		Debug(err)
	} else {
		notification = false
		_, err := db.Exec(`DELETE FROM likes WHERE pid = ? AND uid = ? AND cid = ?`, pid, GetUserByCookies(w, r).Uid, cid)
		Debug(err)
	}

	http.Redirect(w, r, "/", 301)

}
