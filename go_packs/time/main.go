package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	var times int
	for range ticker.C {
		fmt.Println("time.Now:", time.Now())
		times++
		time.Sleep(2 * time.Second)
		fmt.Println("do somethings....")

		if times == 5 {
			fmt.Println("end")
			return
		}
	}
}
