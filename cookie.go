package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter, r *http.Request, uid string) {
	cookie := &http.Cookie{
		Name:       "session",
		Value:      uid,
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     99999999999,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	}
	fmt.Println("Cookie created")
	fmt.Println(cookie)
	http.SetCookie(w, cookie)
}

func GetUserByCookies(w http.ResponseWriter, r *http.Request) User {
	var id int
	var uid string
	var username string
	var email string
	var passwd string

	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Cookie not found")
	} else {
		fmt.Println("Cookie : ", cookie)

		db, _ := sql.Open("sqlite3", "./database.db")

		rows, err := db.Query("SELECT * FROM user WHERE uid='" + cookie.Value + "'")
		Debug(err)

		for rows.Next() {
			err = rows.Scan(&id, &uid, &username, &email, &passwd)
			Debug(err)
		}

	}
	return User{Id: id, Uid: uid, Email: email, Passwd: passwd, Username: username}
}
