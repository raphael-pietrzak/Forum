package forum

import (
	"net/http"
	"text/template"
)

func Filters(w http.ResponseWriter, r *http.Request) {
	Posts2 := []Post{}

	ErrParseForm(w, r)

	category := r.Form.Get("categ")

	if category != "None" {
		for _, v := range Posts {
			if v.Category == category {
				Posts2 = append(Posts2, v)
			}
		}
	} else {
		Posts2 = RecupPost()
	}

	date := r.Form.Get("Date")
	if date == "Decroissante" {
		for i, j := 0, len(Posts2)-1; i < j; i, j = i+1, j-1 {
			Posts2[i], Posts2[j] = Posts2[j], Posts2[i]
		}
	}

	UserLogin := GetUserByCookies(w, r)

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts2, User: UserLogin, PostCategory: Category})
}
