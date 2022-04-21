package forum

import (
	"database/sql"
	"net/http"
	"text/template"
)

func AddPost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		tmpl := template.Must(template.ParseFiles("static/post.html"))
		tmpl.Execute(w, Send{PostCategory: Category})

	case "POST":
		
		db, _ := sql.Open("sqlite3", "./database.db")
		ErrParseForm(w, r)

		_, err := db.Exec(`INSERT INTO posts ('content', 'category', 'uid') VALUES (?, ?, ?)`, r.Form.Get("post_content"), r.Form.Get("Sport"), GetUserByCookies(w, r).Uid)
		Debug(err)

		http.Redirect(w, r, "/", 301)
	}
}
