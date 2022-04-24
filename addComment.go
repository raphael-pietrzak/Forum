package forum

import (
	"database/sql"
	"net/http"
	"time"
)

func AddComment(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	_, err := db.Exec(`INSERT INTO comments ('content', 'pid', 'uid', 'date') VALUES (?, ?, ?, ?)`, r.Form.Get("comment_content"), r.Form.Get("post_id"), GetUserByCookies(w, r).Uid, time.Now().Format("2006-01-02 15:04:05"))
	Debug(err)

	http.Redirect(w, r, "/", 301)
}
