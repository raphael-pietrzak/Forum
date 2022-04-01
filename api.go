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
	Uid    int    `json:"uid"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
	Username string `json:"username"`
}

type Post struct {
	Pid      int    `json:"pid"`
	Category string `json:"category"`
	Picture  string `json:"picture"`
	Comment  []string
}

type Comment struct {
	Cid     int    `json:"cid"`
	Content string `json:"content"`
}

func DataToStruct(w http.ResponseWriter, r *http.Request) {
	// json.Unmarshal(post, &Posts)

	fmt.Fprintf(w, "Hello from API")
}

