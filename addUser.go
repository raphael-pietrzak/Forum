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
		Data.ErrorMessage = ""

	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")

		ErrParseForm(w, r)
		Data.ErrorMessage = ""

		username := r.Form.Get("user_name")
		mail := r.Form.Get("mail")
		password := r.Form.Get("password")
		password_confirm := r.Form.Get("confirm_password")
		uuid := uuid.New()
		avatar := "user.png"
		user := "user"

		if mail == "r@r" {
			avatar = "admin.png"
			user = "admin"
		}

		fmt.Println(mail)
		fmt.Println(password)
		fmt.Println(password_confirm)


		AlreadyUsed := false
		for _, v := range Users {
			fmt.Println(v.Email)
			if v.Email == mail {
				AlreadyUsed = true
				Data.ErrorMessage = "Email déjà utilisé"
			}
		}

		if AlreadyUsed && password != password_confirm {
			Data.ErrorMessage = "Email déjà utilisé et mot de passe incorect"
			http.Redirect(w, r, "/adduser", 301)
		} else if password != password_confirm {
			Data.ErrorMessage = "Not same password"
			http.Redirect(w, r, "/adduser", 301)
		} else {
			if AlreadyUsed  {
				http.Redirect(w, r, "/adduser", 301)
			} else {
				_, err := db.Exec(`INSERT INTO user ('uid', 'username', 'email', 'passwd', 'avatar', 'type') VALUES (?, ?, ?, ?, ?, ?)`, uuid.String(), username, mail, Hash(password), avatar, user)
				Debug(err)
				CreateCookie(w, r, uuid.String())
				http.Redirect(w, r, "/", 301)
			}
		}

	}
}
