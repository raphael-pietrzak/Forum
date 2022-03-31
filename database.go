package forum 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func SqlDatabase(){
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table if not exists user (uid integer not null primary key, username text, email text, passwd text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatal(err)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT IF NOT EXISTS INTO user (username,email,passwd) values('raph','remail','rpasswd'), ('audran','aemail','apasswd'), ('maxim','memail','mpasswd');")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var uid int
		var username string
		var email string 
		var passwd string 
		err = rows.Scan(&uid, &username, &email, &passwd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(uid, username, email, passwd)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}