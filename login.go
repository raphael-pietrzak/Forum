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
		fmt.Println("Users :", Users)

		for i := range Users {
			if Users[i].Email == r.Form.Get("mail") && Users[i].Passwd == r.Form.Get("password") {
				CreateCookie(w, r, Users[i].Uid)
				http.Redirect(w, r, "/", 301)

				
			}
		}

		tmpl := template.Must(template.ParseFiles("static/login.html"))
		tmpl.Execute(w, Posts)
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
		Debug(err)

		http.Redirect(w, r, "/login", 301)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Logout")
	
	// Delete the cookie
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	// http.Redirect(w, r, "/", 301)
}


