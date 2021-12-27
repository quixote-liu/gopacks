package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/hello/https", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, https request")
	}))
	http.ListenAndServe(":8080", nil)
}
