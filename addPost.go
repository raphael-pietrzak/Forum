package forum

import (
	"net/http"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	Posts = append(Posts, Post{Pid: 1, Category: "Categorie 1", Picture: "http://www.google.com/images/srpr/logo11w.png", Comment: []string{"Commentaire 1", "Commentaire 2"}})
	Forum(w, r)
}
