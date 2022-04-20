package main

import (
	"fmt"
	_ "gopacks/config/config"
	"log"

	"github.com/quixote-liu/config"
)

func main() {
	conf := config.CONF

	if err := conf.LoadConfiguration("config.conf"); err != nil {
		log.Println("error:", err)
	}

	fmt.Printf("conf.GetString(\"server\", \"port\"): %v\n", conf.GetString("server", "port"))
	fmt.Printf("conf.GetString(\"server\", \"host\"): %v\n", conf.GetString("server", "host"))
	fmt.Printf("conf.GetString(\"DEFAULT\", \"port\"): %v\n", conf.GetString("DEFAULT", "port"))
	fmt.Printf("conf.GetString(\"DEFAULT\", \"host\"): %v\n", conf.GetString("DEFAULT", "host"))
}
