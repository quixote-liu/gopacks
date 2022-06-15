package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// if err := os.MkdirAll("./file", fs.ModePerm); err != nil {
	// 	log.Println("error:", err)
	// 	return
	// }

	// f, err := os.Create("./file/text.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// if _, err := f.Write([]byte("hello, world")); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// +++++++++++++++++++++++++OPEN+++++++++++++++
	// f, err := os.Open("./file/text.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// dst := make([]byte, 4096)
	// if _, err := f.Read(dst); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println("dst:", string(dst))

	// +++++++++++++++++++++FS+++++++++++++++++++++++++
	fs := os.DirFS("./file")
	f, err := fs.Open("./text.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	dst := make([]byte, 4096)
	if _, err := f.Read(dst); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("dst:", string(dst))
}
