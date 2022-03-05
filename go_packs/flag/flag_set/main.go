package main

import (
	"flag"
	"fmt"
)

func main() {
	start := flag.NewFlagSet("start", flag.ContinueOnError)

	fmt.Println("args:", start.Args())

	start.String("config", "lcs", "this is a name")

	err := start.Parse([]string{"-config", "liuchengshun"})
	if err != nil {
		panic(err)
	}
	fmt.Println("start:", String(start, "config"))

	err = start.Parse([]string{"hahaha", "heiheihei"})
	if err != nil {
		panic(err)
	}
	fmt.Println("args:", start.Args())

	start.Bool("exist", false, "exist or no exist?")
	err = start.Parse([]string{"-exist", "0"})
	if err != nil {
		panic(err)
	}
	fmt.Println("exist:", String(start, "exist"))
}

func String(flag *flag.FlagSet, name string) string {
	return flag.Lookup(name).Value.String()
}
