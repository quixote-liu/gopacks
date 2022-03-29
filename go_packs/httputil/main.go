package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {
	serverA := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("r.Request.Header:", r.Header)

		fmt.Fprintf(rw, "request method: %s url: %s", r.Method, r.URL.String())
	}))

	ua, _ := url.Parse(serverA.URL)
	ua.Path = "/hello/reserve/proxy"

	serverB := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("r.Request.HeaderB:", r.Header)
		proxy := httputil.NewSingleHostReverseProxy(ua)
		proxy.ServeHTTP(rw, r)
	}))

	ub, _ := url.Parse(serverB.URL)
	ub.Path = "/serverB"

	resp, err := http.Get(ub.String())
	if err != nil {
		panic(err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body=", string(b))
}
