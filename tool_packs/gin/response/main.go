package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		err := fmt.Errorf("this is error")

		c.String(http.StatusBadRequest, err.Error())
	})

	r.Run(":8989")
}
