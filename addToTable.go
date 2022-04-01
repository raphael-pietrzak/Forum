package forum 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "log"
)


func AddValues(table string, types string, value string) {
	db, err := sql.Open("sqlite3", "./database.db")
	debug(err)
	execution := "INSERT IF NOT EXISTS INTO " + table + types + " VALUES " + value +";"
	db.Exec(execution);


}

func RecupUser() []User {
	db, err := sql.Open("sqlite3", "./database.db")
	selection := "SELECT * FROM user" 
	rows, err := db.Query(selection)
	debug(err)

	err = rows.Err()
	debug(err)
	var newTab []User
	for rows.Next() {
		var username string 
		var email string 
		var passwd string 
		var uid int
		err = rows.Scan(&uid, &username, &email, &passwd) 
		debug(err)
		newTab = append(newTab, User{ Uid: uid, Email:email, Username:username, Passwd:passwd})
	}
	err = rows.Err()
	debug(err)
	return newTab 
}

