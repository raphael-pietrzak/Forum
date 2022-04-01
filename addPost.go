package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

func AddPost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		tmpl := template.Must(template.ParseFiles("static/post.html"))
		tmpl.Execute(w, Posts)

	case "POST":

		db, _ := sql.Open("sqlite3", "./database.db")

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		post_content := r.Form.Get("post_content")

		_, err := db.Exec("INSERT INTO posts ('content') VALUES ('" + post_content + "');")
		Debug(err)

		Posts = append(Posts, Post{Content: post_content})
		http.Redirect(w, r, "/", 301)

	}
}
