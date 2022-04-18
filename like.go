package forum

import (
	"database/sql"
	"net/http"
	"strings"
)

// func Mike(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("toto")
// 	r.ParseForm()
// 	fmt.Println(r.Form.Get("unlike"))
// 	var var1 string
// 	tmpl := template.Must(template.ParseFiles("static/index.html"))
// 	tmpl.Execute(w, Send{Post: Posts, PostCategory: Category})
// 	var1 = r.Form.Get("unlike")
// 	if var1 == "like" {
// 		fmt.Println("yoyooyoyoyoyo")
// 		db, _ := sql.Open("sqlite3", "./database.db")

// 		ErrParseForm(w, r)

// 		post_content := r.Form.Get("post_content")
// 		post_id, _ := strconv.Atoi(r.Form.Get("post_id"))
// 		categorie := r.Form.Get("Sport")
// 		// like := r.Form.Get("unlike")
// 		UserLogin := GetUserByCookies(w, r)
// 		// fmt.Println(categorie)
// 		SqlExec := `UPDATE posts ('content', 'category', 'uid')
// 		VALUES ('` + post_content + `', '` + categorie + `', '` + UserLogin.Uid + `');`

// 		_, err := db.Exec(SqlExec)
// 		Debug(err)

// 		Posts = append(Posts, Post{Pid: post_id, Content: post_content, Category: categorie, Uid: UserLogin.Uid, User: UserLogin})
// 		http.Redirect(w, r, "/", 301)
// 	}
// 	fmt.Println(var1)
// }

func LikePosts(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./database.db")

	ErrParseForm(w, r)
	likerecup := strings.Split(r.Form.Get("like"), " ")
	like := likerecup[4]
	post_id := likerecup[0]

	

	if like == "like" {
		_, err := db.Exec(`INSERT OR IGNORE INTO likes ('pid', 'uid') VALUES (?, ?)`, post_id, GetUserByCookies(w, r).Uid)
		Debug(err)
	} else {
		_, err := db.Exec(`DELETE FROM likes WHERE pid = ? AND uid = ?`, post_id, GetUserByCookies(w, r).Uid)
		Debug(err)
	}

	// db, _ := sql.Open("sqlite3", "./database.db")
	http.Redirect(w, r, "/", 301)

}
