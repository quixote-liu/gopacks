package main

import "fmt"

func main() {
	str := "hello, world"
	for _, v := range str {
		fmt.Println("string(v):", string(v))

		// vv := []rune{v}
		// string(vv)
	}
}
