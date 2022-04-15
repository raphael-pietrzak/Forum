package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

func Filters(w http.ResponseWriter, r *http.Request) {

	Posts2 := []Post{}

	ErrParseForm(w, r)


	date := r.Form.Get("Date")
	fmt.Println(date)

	category := r.Form.Get("categ")
	fmt.Println(category)
	if category == "None" {
		http.Redirect(w, r, "/", 301)
	} else {
		for _, v := range Posts {
			if v.Category == category {
				Posts2 = append(Posts2, v)
			}
		}
	}
	UserLogin := GetUserByCookies(w, r)

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts2, User: UserLogin, PostCategory: Category})

}
