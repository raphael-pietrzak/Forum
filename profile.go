package forum

import (
	"net/http"
	"text/template"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	MyPosts := []Post{}
	User := GetUserByCookies(w, r)
	for i := range(Posts){
		if Posts[i].Uid == User.Uid{
			MyPosts = append(MyPosts, Posts[i])
		}
	}
	
	UserLogin := GetUserByCookies(w, r)

	tmpl := template.Must(template.ParseFiles("static/profile.html"))
	tmpl.Execute(w, Send{User: UserLogin, Post: MyPosts})
}
