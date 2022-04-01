package forum

import (
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//mail := r.Form.Get("mail")
	//password := r.Form.Get("password")

	tmpl := template.Must(template.ParseFiles("static/login.html"))
	tmpl.Execute(w, Posts)
}

func Sign_up(w http.ResponseWriter, r *http.Request) {
	//username := r.Form.Get("user_name")
	//mail := r.Form.Get("mail")
	//password := r.Form.Get("password")

	tmpl := template.Must(template.ParseFiles("static/sign_up.html"))
	tmpl.Execute(w, Posts)
}
