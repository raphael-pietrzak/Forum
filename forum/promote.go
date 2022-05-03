package forum

import (
	"database/sql"
	"net/http"
	"strings"
)

func Promote(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	infos := r.Form.Get("promotion")
	infoslist := strings.Split(infos, " ")
	promotion := infoslist[0]
	uid := infoslist[1]
	typeUser := infoslist[2]

	if promotion == "promote" {
		if typeUser == "user" {
			_, err := db.Exec(`UPDATE user SET type = ?, avatar = ? WHERE uid = ?`, "modo","modo.png", uid)
			Debug(err)
		} else if typeUser == "modo" {
			_, err := db.Exec(`UPDATE user SET type = ?, avatar = ? WHERE uid = ?`, "admin","admin.png", uid)
			Debug(err)
		}
	} else if promotion == "demote" && uid != GetUserByCookies(w, r).Uid {
		if typeUser == "modo" {
			_, err := db.Exec(`UPDATE user SET type = ?, avatar = ? WHERE uid = ?`, "user","user.png", uid)
			Debug(err)
		} else if typeUser == "admin" {
			_, err := db.Exec(`UPDATE user SET type = ?, avatar = ? WHERE uid = ?`, "modo","modo.png", uid)
			Debug(err)
		}
	}

	http.Redirect(w, r, "/profile", 301)
}
