package forum

import (
	"fmt"
	"net/http"
)

func Filters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// UserLogin := GetUserByCookies(w, r)
		// tmpl := template.Must(template.ParseFiles("static/index.html"))
		// tmpl.Execute(w, Send{Post: Posts, User: UserLogin, PostCategory: Category})
		Home(w, r)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		categ := r.Form.Get("categ")
		fmt.Println(categ)

		Posts2 := Posts
		Posts = []Post{}

		for _, v := range Posts2 {
			if v.Category == categ {
				Posts = append(Posts, v)
			}
		}

		http.Redirect(w, r, "/", 301)
	}
}
