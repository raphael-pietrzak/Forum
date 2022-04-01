package main

import (
	"fmt"
	f "forum"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

var PostContent string

func main() {
	// f.SqlDatabase()

	f.TableCreation()
	fmt.Println(f.RecupUser())

	http.HandleFunc("/", f.Forum)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/addcomment", f.AddComment)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/sign_up", f.Sign_up)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}
