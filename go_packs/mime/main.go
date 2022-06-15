package main

import (
	"fmt"
	"log"
	"mime"
	"path/filepath"
)

func main() {
	{
		path := "/path/file.md"
		res := mime.TypeByExtension(filepath.Ext(path))
		fmt.Printf("res: %v\n", res)
	}

	fmt.Println("=========================")

	{
		path := "/path/file.html"
		var typ string = "text/html; charset=utf-8"
		if err := mime.AddExtensionType(filepath.Ext(path), typ); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("typ: %v\n", typ)
	}

	fmt.Println("=========================")

	{
		ct := mime.QEncoding.Encode("UTF-8", "multipart/mixed")
		fmt.Println("mime.QEncoding.Encode: ", ct)
	}

	{
		ct := mime.TypeByExtension("text")
		fmt.Println("ct:", ct)
		
	}
}
