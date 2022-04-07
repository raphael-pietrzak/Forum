package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Check Cookie
	UserLogin := GetUserByCookies(w, r)

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts, User: UserLogin})
}

func Aaa(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	// //delete cookie
	// // cookie, err := r.Cookie("session")
	// cookie := &http.Cookie{
	// 	Name:   "session",
	// 	Value:  "",
	// 	MaxAge: -1,
	// }
	// // Debug(err)
	// http.SetCookie(w, cookie)

	// http.Redirect(w, r, "/", 301)
}
