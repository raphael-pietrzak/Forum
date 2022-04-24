package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func Login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("static/login.html"))
		tmpl.Execute(w, Posts)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		RecupUser()
		wrong_login := true

		for i := range Users {
			if Users[i].Email == r.Form.Get("mail") && Users[i].Passwd == Hash(r.Form.Get("password")) {
				CreateCookie(w, r, Users[i].Uid)
				wrong_login = false
			}
		}
		if wrong_login {
			http.Redirect(w, r, "/login", 301)
		} else {
			http.Redirect(w, r, "/", 301)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w, r)
	http.Redirect(w, r, "/", 301)
}


func Passwd_forgot(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		
	case "GET":
		tmpl := template.Must(template.ParseFiles("static/passwd_forgot.html"))
		tmpl.Execute(w, Data)
		Data.ErrorMessage = ""
	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")
		ErrParseForm(w, r)
		Data.ErrorMessage = ""

		mail := r.Form.Get("mail")
		password := r.Form.Get("new_password")
		password_confirm := r.Form.Get("confirm_password")

		if CheckMail(mail) == false && password != password_confirm {
			Data.ErrorMessage = "Mail et mot de passe incorrect"
			http.Redirect(w, r, "/forgot", 301)
		} else if password != password_confirm {
			Data.ErrorMessage = "Mot de passe incorrect"
			http.Redirect(w, r, "/forgot", 301)
		} else {
			if CheckMail(mail) == false {
				Data.ErrorMessage = "Mail incorrect"
				http.Redirect(w, r, "/forgot", 301)
			} else {
			_, err := db.Exec(`UPDATE user SET passwd = ? WHERE email = ?`, password, mail)
			Debug(err)
	
			http.Redirect(w, r, "/login", 301)
			}
		}
	}
}



