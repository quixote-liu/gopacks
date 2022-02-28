package main

import "fmt"

func main() {
	var ports []string
	ports = nil

	for _, p := range ports {
		fmt.Println("port:", p)
	}

	fmt.Println("END")
}
