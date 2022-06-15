package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// f, err := os.Create("./file/file.txt")
	f, err := os.Open("./file/file.txt")
	// f, err := os.OpenFile("./file/file.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("create file failed: %v\n", err)
		return
	}
	fmt.Println("file:", f)

	// _, err = f.Write([]byte("hello, world, 2"))
	// if err != nil {
	// 	fmt.Printf("write failed: %v\n", err)
	// 	return
	// }

	dst := bytes.NewBuffer(make([]byte, 4096))
	if _, err = io.Copy(dst, f); err != nil {
		fmt.Printf("copy file connent failed: %v \n", err)
		return
	}
	fmt.Printf("dst.Bytes(): %v\n", dst.String())

	// if _, err = f.Read(dst); err != nil {
	// 	fmt.Printf("read file failed: %v\n", err)
	// 	return
	// }
	// fmt.Println("dst: ", string(dst))
}
