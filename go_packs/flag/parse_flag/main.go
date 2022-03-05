package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "liuchengshun", "this is a name")

func main() {
	flag.Parse()

	fmt.Println("name:", *name)
}
