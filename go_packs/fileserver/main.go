package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/assert/", http.StripPrefix("/assert", http.FileServer(http.Dir("./static"))))
	http.Handle("/health/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "ok")
	}))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
