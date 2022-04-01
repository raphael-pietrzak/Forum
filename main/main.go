package main

import (
	_ "database/sql"
	"fmt"
	f "forum"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// f.SqlDatabase()
	fmt.Println(f.RecupUser()) 

	// db, _ := sql.Open("sqlite3", "./datbase.db")
	// _, err := db.Exec("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, content TEXT, title TEXT);")
	// f.Debug(err)

	http.HandleFunc("/", f.Forum)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/sign_up", f.Sign_up)
	http.HandleFunc("/new_post", f.NewPost)


	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}

