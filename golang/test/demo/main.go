package main

import "fmt"

func main() {
	var n int = 100
	n = n / 1000
	if float64(n) >= 0.1 {
		fmt.Println("hello, float64")
	}
	fmt.Printf("%T, %v\n", n, n)

	str := "hello, world"
	for _, b := range []byte(str) {
		if b > 45 {
			fmt.Printf("%T, %v\n", b, b)
		}
	}
}
