package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	group := r.Group("/go/proxy")

	group.GET("/compute/v2.1/flavors/:flavor_id/os-flavor-access", func(c *gin.Context) {
		flavorID := c.Param("flavor_id")
		c.JSON(http.StatusOK, gin.H{
			"method":    c.Request.Method,
			"hello":     "flavor",
			"flavor_id": flavorID,
		})
	})

	group.GET("/compute/v2.1/flavors/:flavor_id", func(c *gin.Context) {
		flavorID := c.Param("flavor_id")
		c.JSON(http.StatusOK, gin.H{
			"method":    c.Request.Method,
			"hello":     "flavor",
			"flavor_id": flavorID,
		})
	})

	group.Any("/:service/*os_api", func(c *gin.Context) {
		service := c.Param("service")
		osAPI := c.Param("os_api")
		c.JSON(http.StatusOK, gin.H{
			"method":  "any",
			"service": service,
			"os_api":  osAPI,
		})
	})

	r.Run(":8989")
}
