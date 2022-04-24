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
	SumLikes     int
	SumPostLikes int
}

type Post struct {
	User     User
	Uid      string
	Category string
	Content  string
	ContentPhoto string
	Comments []Comment
	Pid      int
	Like 	string
	LikeActive string
	Date       string
}

type Comment struct {
	Uid        string
	Cid        int
	Content    string
	Pid        int
	User       User
	Like       string
	LikeActive string
	Date       string
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

<<<<<<< HEAD

//Utilité Cookie ????
type Cookie struct {
	UUid  string
	Name  string
	Value string
}

type Like struct {
	Uid string
	Pid int
	Uid_receveur string
	Notification_count int
}

=======
>>>>>>> e48cd09cb0b47c4a9a7f1b6c05ac36fd597eb766
func AdminSystem(email string) bool {
	for _, v := range Admin {
		if v == email {
			return true
		}
	}
	return false
}
