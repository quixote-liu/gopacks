package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Host", "127.0.0.1:8989")
	req.Header.Set("scheme", "http")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	payload := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("users:", payload)
}
