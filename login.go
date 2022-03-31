package forum

import (
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/login.html"))
	tmpl.Execute(w, Post{Pid: 1, Category: "Categorie 1", Picture: "http://www.google.com/images/srpr/logo11w.png", Comment: []string{"Commentaire 1", "Commentaire 2"}})
}
