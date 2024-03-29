package main



import (
	"fmt"
	f "forum/forum"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var PostContent string
var Category string

// Main
func main() {

	f.TableCreation()

	f.Category = []string{
		"Sport",
		"Jeux",
		"Nourriture",
	}

	f.Admin = []string{
		"r@r",
		"admin@admin",
		"a@a",
	}

	//HandleFunc
	http.HandleFunc("/", f.Home)
	http.HandleFunc("/profile", f.Profile)

	http.HandleFunc("/adduser", f.AddUser)
	http.HandleFunc("/addpost", f.AddPost)
	http.HandleFunc("/addcomment", f.AddComment)
	http.HandleFunc("/addcategory", f.AddCategory)

	http.HandleFunc("/login", f.Login)
	http.HandleFunc("/logout", f.Logout)
	http.HandleFunc("/forgot", f.Passwd_forgot)

	http.HandleFunc("/avatar", f.ChangementAvatar)
	http.HandleFunc("/like", f.LikePosts)
	http.HandleFunc("/filters", f.Filters)
	http.HandleFunc("/promote", f.Promote)
	http.HandleFunc("/delete", f.Delete)
	http.HandleFunc("/modify", f.Modify)



	//Show #CSS
	fs := http.FileServer(http.Dir("./forum/views"))
	http.Handle("/forum/views/", http.StripPrefix("/forum/views/", fs))

	fs2 := http.FileServer(http.Dir("./forum/profil"))
	http.Handle("/forum/profil/", http.StripPrefix("/forum/profil/", fs2))

	fmt.Println("Listening at http://localhost:8800")
	http.ListenAndServe("localhost:8800", nil)
}

