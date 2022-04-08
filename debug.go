package forum

import (
	"fmt"
	"net/http"
)

func Debug(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ErrParseForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
}
