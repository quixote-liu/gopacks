package main

import "fmt"

func main() {
	r := 18
	nr := fmt.Sprintf("\\U+%04X", int64(r))
	fmt.Println(nr)
}