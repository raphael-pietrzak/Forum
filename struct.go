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
var Category []string

type Send struct {
	Post         []Post
	User         User
	PostCategory []string
}

type Post struct {
	Username string
	Uid      string
	Category string
	Content  string
	Comments []Comment
	Pid      int
}

type Comment struct {
	Uid      string
	Cid      int
	Content  string
	Pid      int
	Username string
}

type User struct {
	Id       int
	Uid      string
	Email    string
	Passwd   string
	Username string
	Avatar   string
	Type     string
}

type Cookie struct {
	UUid  string
	Name  string
	Value string
}

func DataToStruct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API")
}
