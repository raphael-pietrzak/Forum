package main

import (
	"fmt"
	f "forum"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var PostContent string
var Category string


func main() {

	f.TableCreation()

	//Recup Data
	f.Category = []string{"Sport", "Jeux", "Nourriture"}
	f.Users = f.RecupUser()
	f.Posts = f.RecupPost()
	f.Comments = f.RecupComment()

	//HandleFunc
	http.HandleFunc("/", f.Home)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/addcomment", f.AddComment)
	http.HandleFunc("/adduser", f.AddUser)
	http.HandleFunc("/filters", f.Filters)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/forgot", f.Passwd_forgot)
	http.HandleFunc("/logout", f.Logout)
	http.HandleFunc("/profile", f.Profile)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://127.0.0.1:5500")
	http.ListenAndServe("localhost:5500", nil)
}
