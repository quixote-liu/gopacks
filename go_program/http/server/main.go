package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"users": []gin.H{
				{"name": "lcs_testing"},
				{"name": "lcs_testing"},
				{"name": "lcs_testing"},
				{"name": "lcs_testing"},
				{"name": "lcs_testing"},
			},
		})
	})

	r.Run(":8989")
}
