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

	tmpl := template.Must(template.ParseFiles("static/profile.html"))
	tmpl.Execute(w, Send{User: UserLogin, Post: MyPosts})
}

var Avatar string

func ChangementAvatar(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	Avatar = r.Form.Get("image_uploads")
	fmt.Println("avatar")
	fmt.Println(Avatar)

	r.ParseMultipartForm(10 << 20)

	fmt.Println("File Upload Endpoint Hit")
	file, handler, err := r.FormFile("image_uploads")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return 
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	_, err2 := db.Exec("UPDATE user SET avatar = '" + handler.Filename + "' ")
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
