package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Cookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Cookie not found")

	} else {
		fmt.Println("Cookie found")
		fmt.Println("Cookie : ", cookie)

		db, _ := sql.Open("sqlite3", "./database.db")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		rows, err := db.Query("SELECT * FROM user WHERE uid='" + cookie.Value + "'")

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

			UserLogin = User{Id: id, Uid: uid, Email: email, Passwd: passwd, Username: username}

		}
	}
}

