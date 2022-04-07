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
		db, _ := sql.Open("sqlite3", "./database.db")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		mail := r.Form.Get("mail")
		password := r.Form.Get("password")
		fmt.Println(mail, password)

		rows, err := db.Query("SELECT * FROM user WHERE email='" + mail + "' AND passwd='" + password + "'")
		Debug(err)

		for rows.Next() {
			var id int
			var uid string
			var username string
			var email string
			var passwd string
			err = rows.Scan(&id, &uid, &username, &email, &passwd)
			Debug(err)

			fmt.Println("On vous a bien trouvé Monsieur", username)

			// Create a new cookie
			CreateCookie(w, r, uid)

			tmpl := template.Must(template.ParseFiles("static/index.html"))
			tmpl.Execute(w, Send{Post: Posts, User: User{Username: username}})
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
		Debug(err)

		http.Redirect(w, r, "/login", 301)
	}
}



