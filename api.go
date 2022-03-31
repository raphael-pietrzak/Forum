package forum 

import 

struct User {
	Uid int `json:"uid"`
	Email string `json:"email"`
	Passwd string `json:"passwd"`
}

struct Post {
	Pid int `json:"pid"`
	Category string `json:"category"`
	Picture string `json:"picture"`
}

struct Comment {
	Cid int `json:"cid"` 
	Content string `json:"content"`
}
