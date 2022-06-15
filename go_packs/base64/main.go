package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	userPass := "admin@quixote_liu:hubei@lcs_1208"
	dst := make([]byte, base64.StdEncoding.EncodedLen(len([]byte(userPass))))
	base64.StdEncoding.Encode(dst, []byte(userPass))
	fmt.Println("data:", string(dst))
	fmt.Println("data:", len(dst))
}
