package main

import (
	"database/sql"
	"fmt"
	f "forum"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func debug(err error) {
	if err!= nil {
		fmt.Println(err)
	}
}

func main() {
	// f.SqlDatabase()
	fmt.Println(f.RecupUser()) 

	db, _ := sql.Open("sqlite3", "./datbase.db")
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT);")
	debug(err)
	db.Exec("INSERT IF NOT EXISTS INTO people (firstname, lastname) VALUES ('ROB', 'STARK');")
	rows, err  := db.Query("SELECT id, firstname, lastname FROM people;")
	debug(err)
	var id int
	var firstname string
	var lastname string
	for rows.Next(){
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id)+" "+firstname+" "+lastname)
	}
	http.HandleFunc("/", f.Forum)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/sign_up", f.Sign_up)

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)

	
}
