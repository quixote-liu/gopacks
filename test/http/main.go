package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		rw.WriteHeader(http.StatusOK)

		body := map[string]interface{}{
			"name":    "lcs",
			"id":      "8888",
			"address": "hubei",
		}
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		rw.Write(bodyBytes)
	}))
	http.ListenAndServe(":8989", nil)
}
