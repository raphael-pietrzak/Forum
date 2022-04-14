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

		Users = RecupUser()

		for i := range Users {
			if Users[i].Email == r.Form.Get("mail") && Users[i].Passwd == Hash(r.Form.Get("password")) {
				CreateCookie(w, r, Users[i].Uid)
				http.Redirect(w, r, "/", 301)
			}
		}

		tmpl := template.Must(template.ParseFiles("static/login.html"))
		tmpl.Execute(w, Posts)

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
		tmpl.Execute(w, Posts)
	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")
		ErrParseForm(w, r)

		mail := r.Form.Get("mail")
		password := r.Form.Get("new_password")
		fmt.Println(mail, password)

		_, err := db.Exec("UPDATE user SET passwd = '" + password + "' WHERE email='" + mail + "' ")
		Debug(err)

		http.Redirect(w, r, "/login", 301)
	}
}



