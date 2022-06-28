package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.biqugeg.cc/5_5000/490712539.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
