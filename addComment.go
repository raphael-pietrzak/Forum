package forum

import (
	"database/sql"
	"net/http"
)

func AddComment(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	_, err := db.Exec(`INSERT INTO comments ('content', 'pid', 'uid') VALUES (?, ?, ?)`, r.Form.Get("comment_content"), r.Form.Get("post_id"), GetUserByCookies(w, r).Uid)
	Debug(err)

	http.Redirect(w, r, "/", 301)
}
