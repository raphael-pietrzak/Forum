package forum

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"text/template"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	RecupUser()
	RecupPost()
	RecupComment()

	db, err := sql.Open("sqlite3", "./database.db")
	Debug(err)


	
	MyPosts := []Post{}
	User := GetUserByCookies(w, r)
	sumlike := 0
	RecupLike(User)
	_, err2 := db.Exec(`UPDATE likes SET vu = ? WHERE user_liked = ?`, true, User.Uid)
	Debug(err2)




	for i := range Posts {
		if Posts[i].Uid == User.Uid {
			MyPosts = append(MyPosts, Posts[i])
			// sumlike += Posts[i].Like
		}
	}
	


	tmpl := template.Must(template.ParseFiles("static/profile.html"))
	tmpl.Execute(w, Send{User: User, Post: MyPosts, Users: Users, PostCategory: Category, SumLikes: sumlike, Notifications: Data.Notifications, Notifs: Data.Notifs})
}


func ChangementAvatar(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image_uploads")
	Debug(err)

	defer file.Close()

	_, err2 := db.Exec("UPDATE user SET avatar=? WHERE uid=?", handler.Filename, GetUserByCookies(w, r).Uid)
	Debug(err2)

	emptyFile, err := os.Create("profil/" + handler.Filename)
	Debug(err)

	defer emptyFile.Close()

	_, err3 := io.Copy(emptyFile, file)
	Debug(err3)

	http.Redirect(w, r, "/profile", 301)
}



func AddCategory(w http.ResponseWriter, r *http.Request) {
	ErrParseForm(w, r)

	category := r.Form.Get("category")
	Category = append(Category, category)

	http.Redirect(w, r, "/profile", 301)
}

