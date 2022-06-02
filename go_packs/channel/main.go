package main

import "fmt"

func main() {
	TimingMode()
} 

func TimingMode() {
	c := make(chan bool)

	go func(c chan bool) {
		if <-c {
			fmt.Println("in goroutine")
		}

		fmt.Println("END goroutine")
	}(c)

	fmt.Println("precessing")

	c <- true

	fmt.Println("END MAIN")
}
