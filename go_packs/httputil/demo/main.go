package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	s := httptest.NewServer(r)

	r.GET("/hello/reserve/proxy", reserveProxy)
	r.GET("/hello/reserve/proxy/*api", reserveProxy)

	u, _ := url.Parse(s.URL)
	u.Path = "/hello/reserve/proxy"
	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}

	fmt.Println("statuscode:", resp.StatusCode)
}

func reserveProxy(c *gin.Context) {
	if c.Request.URL.Path == "/hello/reserve/proxy" {
		c.Request.Header.Del("X-Forwarded-For")

		fmt.Println("this is hello/reserve/proxy, and doing some things")

		base := &url.URL{}
		base.Host = c.Request.Host
		base.Scheme = "http"
		base.Path = c.Request.RequestURI

		c.Request.URL.Path = "/pragh"
		proxy := httputil.NewSingleHostReverseProxy(base)
		proxy.ServeHTTP(c.Writer, c.Request)
		return
	}

	fmt.Printf("this is %s, and doing some things", c.Request.URL.String())
	c.Status(http.StatusOK)
}
