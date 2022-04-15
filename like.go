package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func Mike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("toto")
	r.ParseForm()
	fmt.Println(r.Form.Get("unlike"))
	var var1 string
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts, PostCategory: Category, Comment: Comments})
	var1 = r.Form.Get("unlike")
	if var1 == "like" {
		fmt.Println("yoyooyoyoyoyo")
		db, _ := sql.Open("sqlite3", "./database.db")

		ErrParseForm(w, r)

		post_content := r.Form.Get("post_content")
		post_id, _ := strconv.Atoi(r.Form.Get("post_id"))
		categorie := r.Form.Get("Sport")
		like := r.Form.Get("unlike")
		UserLogin := GetUserByCookies(w, r)
		// fmt.Println(categorie)
		SqlExec := `UPDATE posts ('content', 'category', 'uid') 
		VALUES ('` + post_content + `', '` + categorie + `', '` + UserLogin.Uid + `');`

		_, err := db.Exec(SqlExec)
		Debug(err)

		Posts = append(Posts, Post{Pid: post_id, Content: post_content, Category: categorie, Uid: UserLogin.Uid, Username: UserLogin.Username})
		http.Redirect(w, r, "/", 301)
	}
	fmt.Println(var1)
}
