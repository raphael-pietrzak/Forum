package forum

import (
	"database/sql"
	"net/http"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	_, err := db.Exec(`DELETE FROM posts WHERE pid = ?`, r.Form.Get("post_id"))
	Debug(err)

	http.Redirect(w, r, "/", 301)
}
