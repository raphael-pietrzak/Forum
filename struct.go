package forum

import (
	_ "github.com/google/uuid"
)

var Admin []string
var Posts []Post
var Users []User
var Comments []Comment
var Data Send
var Category []string

type Send struct {
	Users        []User
	Post         []Post
	User         User
	PostCategory []string
	ErrorMessage string
}

type Post struct {
	User     User
	Uid      string
	Category string
	Content  string
	ContentPhoto string
	Comments []Comment
	Pid      int
	Like 	int
	LikeActive string
}

type Comment struct {
	Uid     string
	Cid     int
	Content string
	Pid     int
	User    User
	Like  	int
	LikeActive string
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


//Utilité Cookie ????
type Cookie struct {
	UUid  string
	Name  string
	Value string
}

type Like struct {
	Uid string
	Pid int
}

func AdminSystem(email string) bool {
	for _, v := range Admin {
		if v == email {
			return true
		}
	}
	return false
}
