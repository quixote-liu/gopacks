package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	info, err := os.Stat("./file.txt")
	// if err != nil {
	// 	log.Println("error: ", err)
	// 	return
	// }
	if err == nil {
		log.Println("exist")
		fmt.Printf("info.Sys(): %v\n", info.Sys())
		return
	}
	fmt.Printf("os.IsExist(err): %v\n", os.IsExist(err))

}
