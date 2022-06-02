package main

import (
	"fmt"
	"net/url"
)

func main() {
	// pathMode()
	pathMode2()
}

func pathMode() {
	path1 := "/api/v1/network"
	path2 := "/api/v1/routers"

	u := url.Values{}
	u.Set("path1", path1)
	u.Set("path2", path2)
	fmt.Printf("url: %s\n", u.Encode())

	p1 := u.Get("path1")
	fmt.Printf("p1: %v\n", p1)
}

func pathMode2() {
	path1 := "/api/v1/network"
	path2 := "/api/v1/routers"

	u := url.Values{}
	u.Add("path1", path1)
	u.Add("path1", path2)
	fmt.Printf("url: %s\n", u.Encode())

	p1 := u.Get("path1")
	fmt.Printf("p1: %v\n", p1)

	p2 := u.Get("path2")
	fmt.Printf("p2: %v\n", p2)
}
