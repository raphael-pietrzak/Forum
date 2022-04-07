package forum

import (
	"fmt"
	"net/http"
)

func Cookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "admin",
		Value: "Admin"}
	for _, c := range r.Cookies() {
		fmt.Println(c.Name, c.Value)
	}

	for _, user := range Users {
		for _, c := range r.Cookies() {
			if c.Value == user.Uid.string() {
				http.Redirect(w, r, "/", 301)
			}
		}
	}

	http.SetCookie(w, &cookie)
}
