package forum

import (
	"fmt"
)

func Debug(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
