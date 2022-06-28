package main

import (
	"os"
	"strings"
	"sync"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fn := "./file"
	// 	data := map[string][]string{
	// 		fn: {strconv.FormatInt(int64(i), 10)},
	// 	}
	// 	go WriteToFile(data, fn)
	// }
	// time.Sleep(3 * time.Second)

	f, err := os.Create("./file")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("hello\nworld")
}

var mu sync.RWMutex

func WriteToFile(content map[string][]string, filename string) {
	mu.Lock()
	defer mu.Unlock()
	msg := marshalContent(content)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(msg)
	if err != nil {
		panic(err)
	}
}

func marshalContent(c map[string][]string) string {
	var cc string
	for k, vv := range c {
		cc += "[" + k + "]" + "\n" + strings.Join(vv, "\n")
	}
	return cc
}
