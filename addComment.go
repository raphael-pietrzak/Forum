package forum

import (
	"database/sql"
	"net/http"
	"strconv"
)

func AddComment(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")

	ErrParseForm(w, r)

	comment_content := r.Form.Get("comment_content")
	id_post, e := strconv.Atoi(r.Form.Get("post_id")) 
	id_post -= 1 
	Debug(e)

	UserLogin := GetUserByCookies(w, r)

	SqlExec := `INSERT INTO comments ('content', 'pid', 'uid') 
	VALUES ('` + comment_content + "','" + r.Form.Get("post_id") + "','" + UserLogin.Uid + "');"


	_, err := db.Exec(SqlExec)
	Debug(err)

	NewComment := Comment{Cid: len(Posts[id_post].Comments) + 1, Content: comment_content, Uid: UserLogin.Uid, Username: UserLogin.Username}

	Posts[id_post].Comments = append(Posts[id_post].Comments, NewComment)
	http.Redirect(w, r, "/", 301)

}
