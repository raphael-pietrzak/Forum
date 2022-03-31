package main

import (
	"fmt"
	f "forum"
	"net/http"
	"golang-sqlite"
)

func main() {
	
	http.HandleFunc("/", f.MainPage)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/view/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}
