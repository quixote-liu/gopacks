package main

import (
	"fmt"
	"sync"
)

type user struct {
	name string
	age  int
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			uu := make([]user, 2)

			var pwg sync.WaitGroup
			for i := range uu {
				pwg.Add(1)
				go func(u *user) {
					defer pwg.Done()
					u.name = "heihei"
				}(&uu[i])

				pwg.Add(1)
				go func(u *user) {
					defer pwg.Done()
					u.age = 2
				}(&uu[i])
			}
			pwg.Wait()

			fmt.Println("users:", uu)
		}()
	}
	wg.Wait()

	fmt.Println("+++++++++++++++END++++++++++++")
}
