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

	pid := os.Getpid()
	fmt.Println("process id:", pid)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("worksplce dirtory: ", wd)

	// err = os.Mkdir("./hello", os.ModeDir)
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.MkdirAll("./hello2/heihei", os.ModeDir)
	// if err != nil {
	// 	panic(err)
	// }

	f, err := os.ReadFile("config")
	if err != nil {
		panic(err)
	}
	fmt.Println("file content:", string(f))
}
