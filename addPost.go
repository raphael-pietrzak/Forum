package forum

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"
)

func AddPost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		tmpl := template.Must(template.ParseFiles("static/post.html"))
		tmpl.Execute(w, Send{PostCategory: Category})

	case "POST":
		
		db, _ := sql.Open("sqlite3", "./database.db")
		ErrParseForm(w, r)

		_, err := db.Exec(`INSERT INTO posts ('content', 'category', 'uid', 'date') VALUES (?, ?, ?, ?)`, r.Form.Get("post_content"), r.Form.Get("Sport"), GetUserByCookies(w, r).Uid, time.Now().Format("2006-01-02 15:04:05"))
		Debug(err)

		http.Redirect(w, r, "/", 301)
	}
}
