package main

import (
	"fmt"
	"net/url"
)

func main() {
	query := "{\"id\": \"d19d04ee-d2e2-4d78-b1cc-cb49cf75f6ff\"}"
	v := url.Values{}
	v.Set("extra_specs", query)

	fmt.Println(v.Encode())
}
