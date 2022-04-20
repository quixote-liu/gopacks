package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Demo2(r)
	r.Run(":8080")
}

func Demo1(r *gin.Engine) {
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
}

func Demo2(r *gin.Engine) {
	r.GET("/*os_api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "gin_router",
		})
	})
	r.GET("/messaging/v2/messages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "messages",
		})
	})

}
