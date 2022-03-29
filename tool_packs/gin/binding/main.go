package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/user", func(c *gin.Context) {
		payload := struct {
			Name  string `json:"name" binding:"required"`
			Age   string `json:"age" binding:"required"`
			Email gin.H  `json:"email" binding:"required"`
		}{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"user": payload,
		})
	})

	r.Run(":8989")
}
