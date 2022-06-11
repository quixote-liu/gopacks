package main

import (
	"fmt"
	"os"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("hostName:", hostName)

	// environ := os.Environ()
	// fmt.Println("environ: \n", environ)

	goRoot := os.Getenv("GOROOT")
	fmt.Println("the go root path:", goRoot)

	goPath := os.Getenv("GOPATH")
	fmt.Println("the go path:", goPath)

	path := os.Getenv("PATH")
	fmt.Println("path:", path)

	fmt.Printf("processIDA(): %v\n", processIDA())
	fmt.Printf("processIDB(): %v\n", processIDB())
}

func processIDA() int {
	return os.Getpid()
}

func processIDB() int {
	return os.Getpid()
}
