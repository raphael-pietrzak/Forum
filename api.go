package forum

import (
	"fmt"
	"net/http"
)

var Posts []Post
var Users []User
var Coment []Comment
var Data Send

type Send struct {
	Post []Post
}

type User struct {
	Uid    int  
	Email  string 
	Passwd string 
	Username string 
}

type Post struct {
	Pid      int   
	Content  string 
	Comment  []string
}

type Comment struct {
	Cid     int    
	Content string 
}

func DataToStruct(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from API")
}

