package forum

import (
	"fmt"
	"net/http"

	_ "github.com/google/uuid"
)

var Posts []Post
var Users []User
var Comments []Comment
var Data Send

type Send struct {
	Post []Post
	User User
}


type Post struct {
	// User User
	Pid      int
	Content  string
	Comments []string
}

type Comment struct {
	Cid     int
	Content string
	Pid    int
}

type User struct {
	Id       int
	Uid      string
	Email    string
	Passwd   string
	Username string
}

type Cookie struct {
	UUid  string
	Name  string
	Value string
}

func DataToStruct(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from API")
	
}
