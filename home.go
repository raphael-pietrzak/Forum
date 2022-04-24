package forum

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {


	UserLogin := GetUserByCookies(w, r)
	RecupUser()
	RecupPost()
	RecupComment()
	RecupLike(UserLogin)

	// Posts = RecupLike(UserLogin.Uid, Posts)

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts, User: UserLogin, PostCategory: Category, Notifications: Data.Notifications, Notifs: Data.Notifs})
}
