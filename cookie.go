package forum

import (
	"fmt"
	"net/http"
	"strconv"
)

func Cookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "admin", 
		Value: "Admin"}
	for _, c := range r.Cookies() {
		fmt.Println(c.Name, c.Value)
	}
	for _, user := range Users {
        for _, c := range r.Cookies() {
            if c.Value == strconv.Itoa(user.Uid) {
				fmt.Println("Cookie found")
				cookie.Value = user.Username
			}
        }
    }


	http.SetCookie(w, &cookie)
}
