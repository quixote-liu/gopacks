package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/api/v1/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, world",
		})
	})

	r.GET("/api/v1/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/hello")
	})

	r.Run(":8989")
}
