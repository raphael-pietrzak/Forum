package forum

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	UserLogin := GetUserByCookies(w, r)

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts, User: UserLogin, PostCategory: Category})
}
