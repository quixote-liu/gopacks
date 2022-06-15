package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	buf := bytes.NewBufferString("heiheihei")

	if err := WriteFile("./file", "text2.txt", buf); err != nil {
		log.Fatal(err)
	}

}

var errFileExist = errors.New("the file is exist")

func WriteFile(path, filename string, read io.Reader) error {
	fn := filepath.Join(path, filepath.Base(filename))
	if _, err := os.Stat(fn); !os.IsNotExist(err) {
		return errFileExist
	}

	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}

	if _, err = io.Copy(f, read); err != nil {
		return fmt.Errorf("write file content failed")
	}

	return nil
}
