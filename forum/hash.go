package forum

import (
	"crypto/sha256"
	"fmt"
)

func Hash(password string) string {
	//hash
	password = password + "salt"
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	password = password[:32]
	return password
}

