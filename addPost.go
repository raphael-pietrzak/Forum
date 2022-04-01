package forum

import (
	"net/http"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	AddValues("user", "('email','passwd')", "('aaaa','bbbbb')")
	Users = RecupUser()
	Forum(w, r)
}
