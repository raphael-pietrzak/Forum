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
		Cookies(w,r)
		tmpl := template.Must(template.ParseFiles("static/login.html"))
		tmpl.Execute(w, Posts)
	case "POST":
		sql.Open("sqlite3", "./database.db")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		mail := r.Form.Get("mail")
		password := r.Form.Get("password")

		Users = RecupUser()

		for _, user := range Users {
			if user.Email == mail && user.Passwd == password {
				http.Redirect(w, r, "/", 301)
			}

		}
		http.Redirect(w, r, "/login", 301)

	}
}

func Passwd_forgot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("static/passwd_forgot.html"))
		tmpl.Execute(w, Posts)
	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		mail := r.Form.Get("mail")
		password := r.Form.Get("new_password")
		fmt.Println(mail, password)

		_, err := db.Exec("UPDATE user SET passwd = '" + password + "' WHERE email='" + mail + "' ")
		debug(err)

		fmt.Println(RecupUser())
		

		http.Redirect(w, r, "/login", 301)
	}
}

func Sign_up(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("static/sign_up.html"))
		tmpl.Execute(w, Posts)
	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		username := r.Form.Get("user_name")
		mail := r.Form.Get("mail")
		password := r.Form.Get("password")

		_, err := db.Exec("INSERT INTO user ('username','email', 'passwd') VALUES ('" + username + "', '" + mail + "', '" + password + "')")
		debug(err)

		http.Redirect(w, r, "/", 301)
	}
}
