package forum

import (
	"database/sql"
	"net/http"
	"strconv"
	"text/template"
)

func AddPost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		tmpl := template.Must(template.ParseFiles("static/post.html"))
		tmpl.Execute(w, Send{PostCategory: Category})

	case "POST":

		db, _ := sql.Open("sqlite3", "./database.db")

		ErrParseForm(w, r)

		post_content := r.Form.Get("post_content")
		post_id,_ := strconv.Atoi(r.Form.Get("post_id"))
		categorie := r.Form.Get("Sport")

		UserLogin := GetUserByCookies(w, r)
		// fmt.Println(categorie)
		SqlExec := `INSERT INTO posts ('content', 'category', 'uid', 'like') 
		VALUES ('` + post_content + `', '` + categorie + `', '` + UserLogin.Uid + `');`

		_, err := db.Exec(SqlExec)
		Debug(err)

		Posts = append(Posts, Post{Pid: post_id, Content: post_content, Category: categorie, Uid: UserLogin.Uid, Username: UserLogin.Username})
		http.Redirect(w, r, "/", 301)
	}
}
