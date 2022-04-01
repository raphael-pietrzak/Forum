package forum

import (
	"fmt"
	"net/http"
	_ "github.com/google/uuid"
)

var Posts []Post
var Users []User
var Coment []Comment
var Data Send

type Send struct {
	Post []Post
}

type User struct {
	Uid      int
	Email    string
	Passwd   string
	Username string
}

type Post struct {
	Pid     int
	Content string
	Comment []string
}

type Comment struct {
	Cid     int
	Content string
}

type Cookie struct {
	UUid  string
	Name  string
	Value string
}

func DataToStruct(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from API")
}
