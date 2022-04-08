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
	f.Category = []string{"Sport", "Jeux", "Nourriture"}

	f.TableCreation()
	f.Posts = f.RecupPost()
	f.Hash("Audran")


	// Cookie

	http.HandleFunc("/", f.Home)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/addcomment", f.AddComment)
	http.HandleFunc("/filters", f.Filters)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/sign_up", f.Sign_up)
	http.HandleFunc("/forgot", f.Passwd_forgot)
	http.HandleFunc("/logout", f.LogOut)
	http.HandleFunc("/profil", f.Profil)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://127.0.0.1:5500")
	http.ListenAndServe("localhost:5500", nil)
}
