package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "strconv"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		DeleteCookie(w, r)
		tmpl := template.Must(template.ParseFiles("static/sign_up.html"))
		tmpl.Execute(w, Data)
		Data.ErrorMessage = ""

	case "POST":
		db, _ := sql.Open("sqlite3", "./database.db")
		ErrParseForm(w, r)
		Data.ErrorMessage = ""

		username := r.Form.Get("user_name") 
		Firstletter := strings.ToUpper(string(username[0])) 
        username = Firstletter + username[1:]
		fmt.Println(username)
		mail := r.Form.Get("mail")
		uuid := uuid.New().String()
		password := r.Form.Get("password")
		password_confirm := r.Form.Get("confirm_password")
		avatar := "user.png"
		user := "user"

		if AdminSystem(mail) {
			avatar = "admin.png"
			user = "admin"
		}

		if CheckMail(mail) && password != password_confirm {
			Data.ErrorMessage = "Email déjà utilisé et mot de passe incorect"
			http.Redirect(w, r, "/adduser", 301)
		} else if password != password_confirm {
			Data.ErrorMessage = "Mot de passe incorrect"
			http.Redirect(w, r, "/adduser", 301)
		} else {
			if CheckMail(mail) {
				http.Redirect(w, r, "/adduser", 301)
			} else {
				_, err := db.Exec(`INSERT INTO user ('uid', 'username', 'email', 'passwd', 'avatar', 'type') VALUES (?, ?, ?, ?, ?, ?)`, uuid, username, mail, Hash(password), avatar, user)
				Debug(err)
				CreateCookie(w, r, uuid)
				http.Redirect(w, r, "/", 301)
			}
		}
	}
}

func CheckMail(mail string) bool {
	for _, v := range Users {
		if v.Email == mail {
			Data.ErrorMessage = "Email déjà utilisé"
			return true
		}
	}
	return false
}
