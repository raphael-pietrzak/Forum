package forum

import (
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("tmpl/index.html"))

	new := rfffr{}
	tmpl.Execute(w, new)
}
