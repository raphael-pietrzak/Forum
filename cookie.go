package forum

import (
	"fmt"
	"net/http"
)

func Cookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "admin", 
		Value: "Admin"}
	for _, c := range r.Cookies() {
		fmt.Println(c.Name, c.Value)
	}


	http.SetCookie(w, &cookie)
}
