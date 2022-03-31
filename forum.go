package forum

import (
	"net/http"
	"text/template"
)

func Forum(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts})
}
