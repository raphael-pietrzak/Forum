package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		tmpl := template.Must(template.ParseFiles("static/sign_up.html"))
		tmpl.Execute(w, Data)

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

		AlreadyUsed := false
		Users := RecupUser()
		for _, v := range Users {
			fmt.Println(v.Email)
			if v.Email == mail {
				AlreadyUsed = true
				Data.ErrorMessage = "Email déjà utilisé"
			}
		}
		if AlreadyUsed == false {
			_, err := db.Exec(`INSERT INTO user ('uid', 'username', 'email', 'passwd', 'avatar', 'type') VALUES (?, ?, ?, ?, ?, ?)`, uuid.String(), username, mail, Hash(password), avatar, user)
			Debug(err)
			CreateCookie(w, r, uuid.String())
			http.Redirect(w, r, "/", 301)
		} else {
			http.Redirect(w, r, "/adduser", 301)
		}

	}
}
