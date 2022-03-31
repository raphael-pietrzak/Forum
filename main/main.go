package main

import (
	"fmt"
	f "forum"
	"net/http"
	"github.com/mattn/go-sqlite3"
)

func main() {
	
<<<<<<< HEAD
	http.HandleFunc("/", f.Forum)
=======
	http.HandleFunc("/", f.MainPage)
>>>>>>> a7a163536db4fced7d1d0934a60fadd35bd2654a

	//Show #CSS
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	fmt.Println("Listening at http://localhost:5500")
	http.ListenAndServe("localhost:5500", nil)
<<<<<<< HEAD
}
=======
}

>>>>>>> a7a163536db4fced7d1d0934a60fadd35bd2654a
