package forum

import (
	"database/sql"
	"net/http"
)

func CreateCookie(w http.ResponseWriter, r *http.Request, uid string) {

	cookie := &http.Cookie{
		Name:  "session",
		Value: uid,
	}
	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

}

func GetUserByCookies(w http.ResponseWriter, r *http.Request) User {

	cookie, err := r.Cookie("session")

	var id int
	var uid string
	var username string
	var email string
	var passwd string
	var avatar string
	typee := "guest"

	if err != nil {
		//fmt.Println("Cookie not found")
	} else {
		//fmt.Println("Cookie found")

		db, _ := sql.Open("sqlite3", "./database.db")

		rows, err := db.Query("SELECT * FROM user WHERE uid='" + cookie.Value + "'")
		Debug(err)

		for rows.Next() {
			err = rows.Scan(&id, &uid, &username, &email, &passwd, &avatar, &typee)
			Debug(err)
		}
	}
	return User{Id: id, Uid: uid, Email: email, Passwd: passwd, Username: username, Avatar: avatar, Type: typee}
}
