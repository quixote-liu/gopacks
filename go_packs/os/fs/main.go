package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fs := os.DirFS("./file")
	f, err := fs.Open("file2.txt")
	if err != nil {
		fmt.Printf("os.IsNotExist(err): %v\n", os.IsNotExist(err))
		log.Fatal(err)
	}
	dst := make([]byte, 4096)
	if _, err := f.Read(dst); err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", string(dst))
}
