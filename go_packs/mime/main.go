package main

import (
	"fmt"
	"mime"
	"path/filepath"
)

func main() {
	path := "/path/file.md"
	res := mime.TypeByExtension(filepath.Ext(path))
	fmt.Printf("res: %v\n", res)
}
