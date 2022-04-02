package forum

import (
	"fmt"
	"net/http"
)

func Cookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Cookie not found")
	} else {
		fmt.Println("Cookie found")
		fmt.Println(cookie)
	}
}
