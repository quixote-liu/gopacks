package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.New()
	gin.SetMode(gin.ReleaseMode)

	c.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "im health",
		})
	})

	c.Run(":8989")
}
