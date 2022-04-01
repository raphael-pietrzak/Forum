package forum

import (
	"net/http"
	"text/template"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	AddValues("user", "('email','passwd')", "('aaaa','bbbbb')")
	Forum(w, r)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	//username := r.Form.Get("user_name")
	//mail := r.Form.Get("mail")
	//password := r.Form.Get("password")

	tmpl := template.Must(template.ParseFiles("static/new_post.html"))
	tmpl.Execute(w, Post{Pid: 1, Category: "Categorie 1", Picture: "http://www.google.com/images/srpr/logo11w.png", Comment: []string{"Commentaire 1", "Commentaire 2"}})
}
