package main

import (
	"fmt"
	"os"
)

func main() {
	root := os.Getenv("GOROOT")
	fmt.Println("go root:", root)

	pid := os.Getpid()
	fmt.Println("pid :", pid)
}
