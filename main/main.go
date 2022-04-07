package main

import (
	"fmt"
	f "forum"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var PostContent string

func main() {

	f.TableCreation()
	f.Posts = f.RecupPost()


	// Cookie

	http.HandleFunc("/", f.Home)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/addcomment", f.AddComment)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/sign_up", f.Sign_up)
	http.HandleFunc("/forgot", f.Passwd_forgot)
	http.HandleFunc("/aaa", f.Aaa)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}
