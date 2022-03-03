package main

import "fmt"

func main() {
	n := new(map[string]string)

	(*n)["service"] = "heihei" // panic

	fmt.Println("service:", (*n)["service"])
}
