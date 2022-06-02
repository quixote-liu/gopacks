package main

import "fmt"

func main() {
	s := "abcdefg"

	var ss string
	for _, v := range s {
		ss += string(v)
	}

	fmt.Printf("ss: %v\n", ss)
}
