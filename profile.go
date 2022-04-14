package forum

import (
	"database/sql"
	"fmt"
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

func ChangementAvatar(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	avatar := r.Form.Get("image_uploads")
	fmt.Println("avatar")
	fmt.Println(avatar)

	_, err := db.Exec("UPDATE user SET avatar = '" + avatar + "' ")
	Debug(err)

	fmt.Println(GetUserByCookies(w,r))
	http.Redirect(w, r, "/profile", 301)
}
