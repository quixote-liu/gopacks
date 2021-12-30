package main

import (
	"fmt"
	"text/template"
)

func main() {
	ts := template.JSEscapeString("function <>hahah")
	fmt.Println(ts)
}