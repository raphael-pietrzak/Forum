package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func Sign_up(w http.ResponseWriter, r *http.Request) {

	// create new uuid

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
		uuid := uuid.New()
		fmt.Println("le nouvel uuid est :", uuid)

		_, err := db.Exec("INSERT INTO user ('uid','username','email', 'passwd') VALUES ('" + uuid.String() + "', '" + username + "', '" + mail + "', '" + password + "')")
		Debug(err)

		CreateCookie(w, r, uuid.String())

		http.Redirect(w, r, "/", 301)

	}
}
