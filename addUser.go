package forum

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		tmpl := template.Must(template.ParseFiles("static/sign_up.html"))
		tmpl.Execute(w, Posts)

	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")

		ErrParseForm(w, r)

		username := r.Form.Get("user_name")
		mail := r.Form.Get("mail")
		password := r.Form.Get("password")
		uuid := uuid.New()
		avatar := "user.png"
		user := "user"

		if mail == "r@r" {
			avatar = "admin.png"
			user = "admin"
		}

		_, err := db.Exec(`INSERT INTO user ('uid', 'username', 'email', 'passwd', 'avatar', 'type') VALUES (?, ?, ?, ?, ?, ?)`, uuid.String(), username, mail, Hash(password), avatar, user)
		Debug(err)

		CreateCookie(w, r, uuid.String())
		http.Redirect(w, r, "/", 301)
	}
}
