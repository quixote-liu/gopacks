package context

import (
	"fmt"
	"log"
	"time"
)

func DoSomethings() {
	fmt.Println("do some things")
	time.Sleep(200 * time.Microsecond)

	go func() {
		fmt.Println("aaa")
		if err := DoWithContext(); err != nil {
			log.Println(err)
			return
		}
	}()

	fmt.Println("End")
}

func DoWithContext() error {
	time.Sleep(500 * time.Microsecond)
	fmt.Println("do in goroutine")
	return nil
}
