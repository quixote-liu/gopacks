package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
)

func main() {
	buf := bytes.NewBufferString("hello, world")
	writer := multipart.NewWriter(buf)
	fmt.Printf("writer.Boundary(): %v\n", writer.Boundary())
	fmt.Printf("writer.Boundary(): %v\n", writer.Boundary())

	fmt.Println("hellodfasdf\r\nworld")
}
