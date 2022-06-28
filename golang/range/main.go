package main

import (
	"fmt"
	"time"
)

func main() {
	var times int
	ticker := time.NewTicker(500 * time.Microsecond)
	for range ticker.C {
		if times == 7 {
			return
		}
		fmt.Printf("times: %v\n", times)
		times++

	}
}
