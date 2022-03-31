package main

import (
	"fmt"
	f "forum"
	"net/http"
	//"github.com/mattn/go-sqlite3"
)

func main() {
	
	http.HandleFunc("/", f.Forum)
	http.HandleFunc("/login", f.Login)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}
