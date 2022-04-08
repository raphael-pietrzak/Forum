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
		tmpl.Execute(w, Send{PostCategory: Category})

	case "POST":

		db, _ := sql.Open("sqlite3", "./database.db")

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		post_content := r.Form.Get("post_content")
		categorie := r.Form.Get("Sport")
		fmt.Println(categorie)


		_, err := db.Exec("INSERT INTO posts ('content') VALUES ('" + post_content + "');")
		Debug(err)

		Posts = append(Posts, Post{Pid: len(Posts) + 1 , Content: post_content, Category: categorie})
		http.Redirect(w, r, "/", 301)
	}
}
