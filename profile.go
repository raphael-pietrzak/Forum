package forum

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	MyPosts := []Post{}
	User := GetUserByCookies(w, r)
	for i := range Posts {
		if Posts[i].Uid == User.Uid {
			MyPosts = append(MyPosts, Posts[i])
		}
	}

	UserLogin := GetUserByCookies(w, r)
	Users = RecupUser()

	tmpl := template.Must(template.ParseFiles("static/profile.html"))
	tmpl.Execute(w, Send{User: UserLogin, Post: MyPosts, Users: Users, PostCategory: Category})
}

var Avatar string

func ChangementAvatar(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	Avatar = r.Form.Get("image_uploads")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image_uploads")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return 
	}
	defer file.Close()
	User := GetUserByCookies(w, r)

	_, err2 := db.Exec("UPDATE user SET avatar=? WHERE uid=?", handler.Filename, User.Uid)
	Debug(err2)

	emptyFile, err := os.Create("profil/" + handler.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer emptyFile.Close()

	_, err3 := io.Copy(emptyFile, file)
	Debug(err3)

	http.Redirect(w, r, "/profile", 301)
}

func AddCategory(w http.ResponseWriter, r *http.Request){
	ErrParseForm(w, r)
	category := r.Form.Get("category")
	Category = append(Category, category)
	http.Redirect(w, r, "/profile", 301)
}