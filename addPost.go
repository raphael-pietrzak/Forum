package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	Posts = []
	db, _ := sql.Open("sqlite3", "./datbase.db")
	db.Exec("INSERT IF NOT EXISTS INTO posts (content, title) VALUES ('ROB', 'STARK');")
	rows, err := db.Query("SELECT * FROM posts;")
	Debug(err)
	var pid int
	var content string
	var title string
	for rows.Next() {
		rows.Scan(&pid, &content, &title)
		fmt.Println(pid, content, title)
		Posts = append(Posts, )
	}

	Forum(w, r)
}
