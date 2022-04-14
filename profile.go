package forum

import (
	"net/http"
	"text/template"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	
	UserLogin := GetUserByCookies(w, r)

	tmpl := template.Must(template.ParseFiles("static/profile.html"))
	tmpl.Execute(w, Send{User: UserLogin})
}
