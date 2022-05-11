package main

import (
	"fmt"
	"net/http"
)

func main() {
	HeaderAddSet()
}

func HeaderAddSet() {
	header1 := http.Header{}

	header1.Add("Content-Type", "application/json")
	header1.Add("Accept-Language", "zh-CN")
	fmt.Println("header:", header1)
}
