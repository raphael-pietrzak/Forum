package forum

import (
	"net/http"
	"sort"
	"text/template"
)

func Filters(w http.ResponseWriter, r *http.Request) {
	Posts2 := []Post{}

	ErrParseForm(w, r)
	UserLogin := GetUserByCookies(w, r)

	category := r.Form.Get("categ")

	if category != "None" {
		for _, v := range Posts {
			if v.Category == category {
				Posts2 = append(Posts2, v)
			}
		}
	} else {
		RecupPost()
		RecupComment()
		RecupLike(UserLogin)
		Posts2 = Posts
	}

	date := r.Form.Get("Date")
	if date == "Decroissante" {
		for i, j := 0, len(Posts2)-1; i < j; i, j = i+1, j-1 {
			Posts2[i], Posts2[j] = Posts2[j], Posts2[i]
		}
	}

	likes := r.Form.Get("LikesFiltres")
	if likes == "plikes" {
		sort.SliceStable(Posts2, func(i, j int) bool {
			return Posts2[i].Like2 < Posts2[j].Like2
		})
	}
	if likes == "mlikes" {
		sort.SliceStable(Posts2, func(i, j int) bool {
			return Posts2[i].Like2 > Posts2[j].Like2
		})
	}
	

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, Send{Post: Posts2, User: UserLogin, PostCategory: Category})
}
