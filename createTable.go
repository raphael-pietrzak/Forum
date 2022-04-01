package forum 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func debug(err error) {
	if err!= nil {
		fmt.Println(err)
	}
}

func tableCreation(){
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (uid INTEGER PRIMARY KEY, username TEXT, email TEXT, passwd TEXT);")
	if err != nil {
		debug(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (pid INTEGER, photo BLOB, category TEXT, FOREIGN KEY(pid) REFERENCES user(uid);")
	if err != nil {
		debug(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS comments (cid INTEGER, content TEXT, FOREIGN KEY(cid) REFERENCES user(uid));")
	if err != nil {
		debug(err)
	}
}
