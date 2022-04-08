package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	UserLogin := GetUserByCookies(w, r)

	categorie := r.Form.Get("Sport")
	fmt.Println(categorie)
	

	tmpl := template.Must(template.ParseFiles("static/profil.html"))
	tmpl.Execute(w, Send{User: UserLogin})
}
