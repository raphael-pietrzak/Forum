package forum

import (
	"net/http"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	AddValues("user", "('email','passwd')", "('aaaa','bbbbb')")
	Posts = RecupPosts()
	Forum(w, r)
}
