package forum

import (
	"fmt"
	"net/http"
)




type User struct {
	Uid    int    `json:"uid"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
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

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API")
}
