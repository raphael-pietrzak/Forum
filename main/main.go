package main

import (
	"fmt"
	f "forum"
	"net/http"
)

func main() {
	http.HandleFunc("/accueil", f.MainPage())

	//Show #CSS
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
}
