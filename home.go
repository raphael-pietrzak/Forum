package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	UserLogin := GetUserByCookies(w, r)

	fmt.Println(Posts)
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts, User: UserLogin, PostCategory: Category})
}
