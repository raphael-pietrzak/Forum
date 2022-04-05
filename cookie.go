package forum

import (
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

		// db, _ := sql.Open("sqlite3", "./database.db")
		// if err := r.ParseForm(); err != nil {
		// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
		// 	return
		// }

		// rows, err := db.Query("SELECT * FROM user WHERE uid='" + cookie.Value + "'")
		// Debug(err)

		// for rows.Next() {
		// 	var id int
		// 	var uid string
		// 	var username string
		// 	var email string
		// 	var passwd string
		// 	err = rows.Scan(&id, &uid, &username, &email, &passwd)
		// 	Debug(err)

		// 	fmt.Println("On vous a bien trouvé Monsieur", username)
		// 	fmt.Println("Votre id est :", id)
		// 	fmt.Println("Votre uid est :", uid)
		// 	fmt.Println("Votre mail est :", email)
		// 	fmt.Println("Votre passwd est :", passwd)

		// 	tmpl := template.Must(template.ParseFiles("static/index.html"))
		// 	tmpl.Execute(w, Send{Post: Posts, User: User{Username: username}})
		// }
	}
}

func Create_Cookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:  "session",
		Value: "cookie-value",
	}

	// Create a new cookie
	fmt.Println("Cookie created")
	fmt.Println(cookie)
	http.SetCookie(w, cookie)
}
