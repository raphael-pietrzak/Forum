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
	Notifications []string
	Notifs int
}

type Post struct {
	User     User
	Uid      string
	Category string
	Content  string
	ContentPhoto string
	Comments []Comment
	Pid      int
	Like2     int
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

func AdminSystem(email string) bool {
	for _, v := range Admin {
		if v == email {
			return true
		}
	}
	return false
}
