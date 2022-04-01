package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//mail := r.Form.Get("mail")
	//password := r.Form.Get("password")

	tmpl := template.Must(template.ParseFiles("static/login.html"))
	tmpl.Execute(w, Posts)
}

func Sign_up(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("static/sign_up.html"))
		tmpl.Execute(w, Posts)
	case "POST":
		db, _ := sql.Open("sqlite3","./database.db")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		username := r.Form.Get("user_name")
		mail := r.Form.Get("mail")
		password := r.Form.Get("password")

		fmt.Println(username)


		_,err := db.Exec("INSERT INTO user ('username','email', 'passwd') VALUES ('"+username+"', '"+mail+"', '"+password+"')")
		debug(err)
		
		fmt.Println(RecupUser())

		http.Redirect(w, r, "/", 301)
	}
}
