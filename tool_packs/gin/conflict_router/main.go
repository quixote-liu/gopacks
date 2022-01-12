package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:service/*os_api", func(c *gin.Context) {
		service := c.Param("service")
		osAPI := c.Param("os_api")

		c.JSON(http.StatusOK, gin.H{
			"path":    c.Request.URL.Path,
			"method":  c.Request.Method,
			"service": service,
			"os_api":  osAPI,
		})
	})

	// r.GET("/:hello/*world", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"hello": "world",
	// 	})
	// })
	r.Run(":8080")
}
