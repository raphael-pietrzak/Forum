package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func AddComment(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	comment_content := r.Form.Get("comment_content")
	id_post, e := strconv.Atoi(r.Form.Get("post_id")) 
	id_post = id_post - 1
	Debug(e)


	_, err := db.Exec("INSERT INTO comments ('content') VALUES ('" + comment_content + "');")
	Debug(err)

	Posts[id_post].Comments = append(Posts[id_post].Comments, comment_content)
	http.Redirect(w, r, "/", 301)
}
