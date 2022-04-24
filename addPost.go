package forum

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

func AddPost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		tmpl := template.Must(template.ParseFiles("static/post.html"))
		tmpl.Execute(w, Send{PostCategory: Category})

	case "POST":

		db, _ := sql.Open("sqlite3", "./database.db")
		ErrParseForm(w, r)

		post := ""

		r.ParseMultipartForm(10 << 20)

		post_content := r.Form.Get("post_content")
		file, handler, err1 := r.FormFile("post-image")
		Debug(err1)

		if handler != nil {
			post = handler.Filename
			emptyFile, err2 := os.Create("profil/" + handler.Filename)
			Debug(err2)

			defer emptyFile.Close()

			_, err3 := io.Copy(emptyFile, file)
			Debug(err3)
		} else {
			post = "rien"
		}

		_, err := db.Exec(`INSERT INTO posts (content, contentPhoto, category, uid, date) VALUES (?, ?, ?, ?, ?)`, post_content, post, r.Form.Get("Sport"), GetUserByCookies(w, r).Uid, time.Now().Format("2006-01-02 15:04:05"))
		Debug(err)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	

	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)

	_, err := db.Exec(`DELETE FROM posts WHERE pid = ?`, r.Form.Get("pid"))
	Debug(err)

	_, err2 := db.Exec(`DELETE FROM comments WHERE pid = ?`, r.Form.Get("pid"))
	Debug(err2)


	http.Redirect(w, r, "/profile", 301)
}

func Modify(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("sqlite3", "./database.db")
	ErrParseForm(w, r)
	if r.Form.Get("post_content") != "" {
		_, err := db.Exec(`UPDATE posts SET content = ? WHERE pid = ?`, r.Form.Get("post_content"), r.Form.Get("pid"))
		Debug(err)
	}
	

	http.Redirect(w, r, "/profile", 301)

}